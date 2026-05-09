package game

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/adm87/onyx/cmd/onyx-game/internal/game/input/bindings"
	"github.com/adm87/onyx/cmd/onyx-game/internal/game/scenes/gameplay"
	"github.com/adm87/onyx/cmd/onyx-game/internal/game/scenes/splashscreen"
	"github.com/adm87/onyx/pkg/engine"
	"github.com/adm87/onyx/pkg/images"
	"github.com/hajimehoshi/ebiten/v2"
)

func Boot(cfg *engine.Config) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	logger := createLogger()
	assets := createAssets(logger)
	input := createInput(logger)
	screen := createScreen(cfg, logger)
	renderer := createRenderer(logger)

	scenes := createScenes(logger, assets, input, screen, renderer)
	if err := scenes.Start(engine.SceneId(cfg.InitialScene)); err != nil {
		logger.Error("failed starting initial scene", "error", err)
		return err
	}

	onyx := newGame(ctx, cfg, logger, screen, scenes)
	if err := engine.Run(cfg, onyx); err != nil {
		logger.Error("game loop exited with error", "error", err.Error())
		return err
	}

	logger.Debug("game loop exited, shutting down...")
	return nil
}

func createLogger() *engine.Logger {
	logger := engine.NewLogger()
	logger.SetLevel(engine.LogLevelDebug)
	logger.Debug("booting game...")
	return logger
}

func createAssets(logger *engine.Logger) *engine.Assets {
	assets := engine.NewAssets(logger)
	assets.RegisterAdapter(images.NewEbitenImageAdapter(logger))
	return assets
}

func createInput(logger *engine.Logger) *engine.Input {
	input := engine.NewInput(logger)
	input.Bind(
		bindings.NewGameQuitBinding(),
		bindings.NewFullscreenToggleBinding(),
	)
	return input
}

func createScreen(cfg *engine.Config, logger *engine.Logger) *engine.Screen {
	screen := engine.NewScreen(
		cfg.Width,
		cfg.Height,
		ebiten.FilterPixelated,
		engine.ScreenResizeByHeight,
	)
	return screen
}

func createRenderer(logger *engine.Logger) *engine.Renderer {
	renderer := engine.NewRenderer(logger)
	return renderer
}

func createScenes(logger *engine.Logger, assets *engine.Assets, input *engine.Input, screen *engine.Screen, renderer *engine.Renderer) *engine.Scenes {
	scenes := engine.NewScenes(logger)
	scenes.Add(
		&engine.SceneEntry{
			Id: splashscreen.SceneId,
			Ctor: func() engine.Scene {
				return splashscreen.NewScene(logger, assets, screen)
			},
			Transitions: map[engine.SceneExitCode]engine.SceneId{
				splashscreen.Complete: gameplay.SceneId,
			},
		},
		&engine.SceneEntry{
			Id: gameplay.SceneId,
			Ctor: func() engine.Scene {
				return gameplay.NewScene(logger, assets, input, screen)
			},
		},
	)
	return scenes
}
