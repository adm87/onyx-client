package images

import (
	"bytes"

	"github.com/adm87/onyx/pkg/engine"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var adapterID engine.AssetAdapterID = "ebiten_image_adapter"

type ImageAssets struct {
	logger *engine.Logger
	cache  *cache
}

func NewEbitenImageAdapter(logger *engine.Logger) *ImageAssets {
	return &ImageAssets{
		logger: logger,
		cache:  newCache(),
	}
}

func (a *ImageAssets) SupportedTypes() []engine.FileType {
	return []engine.FileType{"png", "jpeg", "jpg"}
}

func (a *ImageAssets) ID() engine.AssetAdapterID {
	return adapterID
}

func (a *ImageAssets) Import(path engine.FilePath, data []byte) error {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(data))
	if err != nil {
		a.logger.Error("failed to import image asset", "path", path, "error", err.Error())
		return err
	}
	a.cache.Set(path, img)
	return nil
}

func (a *ImageAssets) Delete(path engine.FilePath) error {
	a.cache.Delete(path)
	return nil
}
