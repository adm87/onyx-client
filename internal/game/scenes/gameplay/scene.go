package gameplay

import (
	"github.com/adm87/onyx/internal/game/input/bindings"
	"github.com/adm87/onyx/pkg/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	SceneId engine.SceneId = "gameplay"
)

type Scene struct {
	logger *engine.Logger
	assets *engine.Assets
	input  *engine.Input
	screen *engine.Screen
}

func NewScene(logger *engine.Logger, assets *engine.Assets, input *engine.Input, screen *engine.Screen) *Scene {
	return &Scene{
		logger: logger,
		assets: assets,
		input:  input,
		screen: screen,
	}
}

func (s *Scene) OnEnter() error {
	s.input.EnableBinding(bindings.Quit)
	s.input.EnableBinding(bindings.Fullscreen)
	return nil
}

func (s *Scene) OnExit() error {
	s.input.DisableBinding(bindings.Quit)
	s.input.DisableBinding(bindings.Fullscreen)
	return nil
}

func (s *Scene) Update() (engine.SceneExitCode, error) {
	if err := s.input.Poll(); err != nil {
		return engine.SceneExitCodeNone, err
	}
	return engine.SceneExitCodeNone, nil
}

func (s *Scene) Draw(screen *ebiten.Image) error {
	safeArea := s.screen.SafeArea()
	minX, minY := safeArea.Min()

	ebitenutil.DebugPrintAt(screen, "Gameplay Scene", int(minX), int(minY))
	return nil
}
