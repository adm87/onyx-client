package gameplay

import (
	"github.com/adm87/onyx/internal/content"
	"github.com/adm87/onyx/internal/ecs/transform"
	"github.com/adm87/onyx/internal/game/input/bindings"
	"github.com/adm87/onyx/pkg/engine"
	"github.com/adm87/onyx/pkg/images"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
)

const (
	SceneId engine.SceneId = "gameplay"
)

type Scene struct {
	logger *engine.Logger
	assets *engine.Assets
	input  *engine.Input
	screen *engine.Screen

	world      donburi.World
	testEntity donburi.Entity
	testImg    *ebiten.Image

	transformSync *transform.TransformSyncSystem
}

func NewScene(logger *engine.Logger, assets *engine.Assets, input *engine.Input, screen *engine.Screen) *Scene {
	world := donburi.NewWorld()
	syncSys := transform.NewTransformSyncSystem()
	return &Scene{
		logger:        logger,
		assets:        assets,
		input:         input,
		screen:        screen,
		world:         world,
		transformSync: syncSys,
	}
}

func (s *Scene) OnEnter() error {
	s.input.EnableBinding(bindings.Quit)
	s.input.EnableBinding(bindings.Fullscreen)

	cache, err := images.Cache(s.assets)
	if err != nil {
		s.logger.Error("failed to get image cache", "error", err)
		return nil
	}

	img, exists := cache.Get(content.StatisImg10x10FilePath)
	if !exists {
		s.logger.Error("failed to get test image from cache", "path", "assets/test.png")
		return nil
	}
	s.testImg = img

	s.testEntity = s.world.Create(
		transform.Transform,
	)

	safeArea := s.screen.SafeArea()
	minX, minY := safeArea.Min()
	maxX, maxY := safeArea.Max()

	entry := s.world.Entry(s.testEntity)
	transform.SetPositionXY(entry, (minX+maxX)/2, (minY+maxY)/2)
	return nil
}

func (s *Scene) OnExit() error {
	s.input.DisableBinding(bindings.Quit)
	s.input.DisableBinding(bindings.Fullscreen)
	return nil
}

func (s *Scene) Update(deltaTime, fixedTime float64, steps int) (engine.SceneExitCode, error) {
	if err := s.input.Poll(); err != nil {
		return engine.SceneExitCodeNone, err
	}

	s.transformSync.Update(s.world)

	return engine.SceneExitCodeNone, nil
}

func (s *Scene) Draw(screen *ebiten.Image) error {
	safeArea := s.screen.SafeArea()
	minX, minY := safeArea.Min()

	entry := s.world.Entry(s.testEntity)
	screen.DrawImage(s.testImg, &ebiten.DrawImageOptions{
		GeoM: *transform.GetGeoM(entry),
	})

	ebitenutil.DebugPrintAt(screen, "Gameplay Scene", int(minX), int(minY))
	return nil
}
