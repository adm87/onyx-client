package splashscreen

import (
	"github.com/adm87/onyx/internal/content"
	"github.com/adm87/onyx/pkg/engine"
	"github.com/adm87/onyx/pkg/images"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

const (
	Complete engine.SceneExitCode = iota + 1
)

const (
	SceneId engine.SceneId = "splash_screen"
)

type Scene struct {
	logger *engine.Logger
	assets *engine.Assets
	screen *engine.Screen
	img    *ebiten.Image
	seq    *gween.Sequence

	complete bool
	opacity  float32
}

func NewScene(logger *engine.Logger, assets *engine.Assets, screen *engine.Screen) *Scene {
	return &Scene{
		logger: logger,
		assets: assets,
		screen: screen,
		seq: gween.NewSequence(
			gween.New(0.0, 1.0, 0.5, ease.Linear),
			gween.New(1.0, 1.0, 2.0, ease.Linear),
			gween.New(1.0, 0.0, 0.5, ease.Linear),
		),
	}
}

func (s *Scene) OnEnter() error {
	// The splash screen errors should be non-blocking
	if err := s.assets.Load(content.StaticAssetFS, content.StaticSplashScreenFilePath); err != nil {
		s.logger.Error("failed to load splash screen", "error", err)
	} else if cache, err := images.Cache(s.assets); err != nil {
		s.logger.Error("failed to get image cache", "error", err)
	} else if img, exists := cache.Get(content.StaticSplashScreenFilePath); !exists {
		s.logger.Error("failed to get splash screen image", "path", content.StaticSplashScreenFilePath)
	} else {
		s.screen.ResizeBuffer(img.Bounds().Dx(), img.Bounds().Dy())
		s.img = img
	}
	return nil
}

func (s *Scene) OnExit() error {
	s.assets.Unload(content.StaticSplashScreenFilePath)
	s.screen.RestoreBuffer()
	return nil
}

func (s *Scene) Update(deltaTime, fixedTime float64, steps int) (engine.SceneExitCode, error) {
	o, _, complete := s.seq.Update(float32(deltaTime))
	s.opacity = o

	if complete {
		return Complete, nil
	}
	return engine.SceneExitCodeNone, nil
}

func (s *Scene) Draw(screen *ebiten.Image) error {
	if s.img == nil {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(float32(s.opacity))
	screen.DrawImage(s.img, op)
	return nil
}
