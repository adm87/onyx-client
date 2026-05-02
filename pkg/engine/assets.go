package engine

import (
	"fmt"
	"io/fs"
)

type AssetAdapterID string

type AssetAdapter interface {
	Import(path FilePath, data []byte) error
	Delete(path FilePath) error

	ID() AssetAdapterID
	SupportedTypes() []FileType
}

type AssetCache[T any] interface {
	Get(path FilePath) (T, bool)
	Set(path FilePath, asset T)
}

type Assets struct {
	logger *Logger

	adaptersByType map[FileType]AssetAdapter
	adaptersByID   map[AssetAdapterID]AssetAdapter
}

func NewAssets(logger *Logger) *Assets {
	return &Assets{
		logger:         logger,
		adaptersByType: make(map[FileType]AssetAdapter),
		adaptersByID:   make(map[AssetAdapterID]AssetAdapter),
	}
}

func (s *Assets) RegisterAdapter(adapter AssetAdapter) {
	if _, exists := s.adaptersByID[adapter.ID()]; exists {
		s.logger.Warn("attempted to register asset adapter with duplicate ID", "adapter_id", adapter.ID())
		return
	}

	for _, t := range adapter.SupportedTypes() {
		if existing, exists := s.adaptersByType[t]; exists {
			s.logger.Warn("attempted to register asset adapter for type already supported by another adapter",
				"adapter_id", adapter.ID(),
				"asset_type", t,
				"existing_adapter_id", existing.ID(),
			)
			continue
		}
		s.adaptersByType[t] = adapter
		s.logger.Debug("registered asset adapter for type", "adapter_id", adapter.ID(), "asset_type", t)
	}

	s.adaptersByID[adapter.ID()] = adapter
	s.logger.Debug("registered asset adapter", "adapter_id", adapter.ID())
}

func (s *Assets) Load(filesystem fs.FS, filepaths ...FilePath) error {
	for _, filepath := range filepaths {
		ftype := filepath.Type()
		if ftype.IsEmpty() {
			return fmt.Errorf("cannot infer asset type from path", "path", filepath)
		}

		adapter, exists := s.adaptersByType[ftype]
		if !exists {
			return fmt.Errorf("unsupported asset type", "type", ftype)
		}

		data, err := fs.ReadFile(filesystem, filepath.String())
		if err != nil {
			return fmt.Errorf("failed to read asset", "error", err)
		}

		if err := adapter.Import(filepath, data); err != nil {
			return fmt.Errorf("failed to import asset", "error", err)
		}
	}
	return nil
}

func (s *Assets) Unload(filepaths ...FilePath) {

}

func (s *Assets) AdapterByType(t FileType) (AssetAdapter, bool) {
	adapter, exists := s.adaptersByType[t]
	return adapter, exists
}

func (s *Assets) AdapterByID(id AssetAdapterID) (AssetAdapter, bool) {
	adapter, exists := s.adaptersByID[id]
	return adapter, exists
}
