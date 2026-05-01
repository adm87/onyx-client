package splashscreen

import (
	"github.com/adm87/onyx/internal/game/input/bindings"
	"github.com/adm87/onyx/pkg/encoding"
	"github.com/adm87/onyx/pkg/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var SplashScreenSceneId = engine.SceneID(encoding.TypeID[Scene]())

type Scene struct {
	logger *engine.Logger
	assets *engine.Assets
	screen *engine.Screen
	input  *engine.Input
}

func New(logger *engine.Logger, assets *engine.Assets, screen *engine.Screen, input *engine.Input) *Scene {
	return &Scene{
		logger: logger,
		assets: assets,
		screen: screen,
		input:  input,
	}
}

func (s *Scene) OnEnter() error {
	s.input.EnableBinding(bindings.QuitBindingID)
	s.input.EnableBinding(bindings.FullscreenToggleBindingID)
	return nil
}

func (s *Scene) OnExit() error {

	return nil
}

func (s *Scene) Update() error {

	return nil
}

func (s *Scene) Draw(screen *ebiten.Image) error {
	safeArea := s.screen.SafeArea()
	minX, minY := safeArea.Min()

	ebitenutil.DebugPrintAt(screen, "Hello, Onyx!", int(minX), int(minY))
	return nil
}
