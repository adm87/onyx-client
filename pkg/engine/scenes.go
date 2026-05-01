package engine

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type SceneID uint8

type Scene interface {
	OnEnter() error
	OnExit() error
	Update() error
	Draw(screen *ebiten.Image) error
}

type SceneEntry struct {
	Id   SceneID
	Ctor func() Scene
}

type Scenes struct {
	logger *Logger

	current Scene
	scenes  map[SceneID]*SceneEntry
}

func NewScenes(logger *Logger) *Scenes {
	return &Scenes{
		logger: logger,
		scenes: make(map[SceneID]*SceneEntry),
	}
}

func (s *Scenes) AddScenes(scenes ...*SceneEntry) {
	for _, entry := range scenes {
		if _, ok := s.scenes[entry.Id]; ok {
			s.logger.Warn("scene %d already exists, skipping", entry.Id)
			continue
		}
		s.scenes[entry.Id] = entry
	}
}

func (s *Scenes) Start(id SceneID) error {
	if err := s.Exit(); err != nil {
		return err
	}

	entry, ok := s.scenes[id]
	if !ok {
		return fmt.Errorf("scene %d not found", id)
	}

	s.current = entry.Ctor()
	return s.current.OnEnter()
}

func (s *Scenes) Exit() error {
	if s.current != nil {
		return s.current.OnExit()
	}
	return nil
}

func (s *Scenes) Update() error {
	if s.current != nil {
		return s.current.Update()
	}
	return nil
}

func (s *Scenes) Draw(screen *ebiten.Image) error {
	if s.current != nil {
		return s.current.Draw(screen)
	}
	return nil
}
