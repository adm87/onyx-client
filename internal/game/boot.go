package game

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/adm87/onyx/internal/game/input/bindings"
	"github.com/adm87/onyx/internal/game/scenes/splashscreen"
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
	scenes, id := createScenes(logger, assets, input, screen)

	if err := scenes.Start(id); err != nil {
		logger.Error("failed starting initial scene: %v", err)
		return err
	}

	onyx := newGame(ctx, cfg, logger, input, assets, screen, scenes)
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
		ebiten.FilterLinear,
		engine.ScreenResizeByHeight,
	)
	return screen
}

func createScenes(logger *engine.Logger, assets *engine.Assets, input *engine.Input, screen *engine.Screen) (*engine.Scenes, engine.SceneID) {
	scenes := engine.NewScenes(logger)
	scenes.AddScenes(
		&engine.SceneEntry{
			Id: splashscreen.SplashScreenSceneId,
			Ctor: func() engine.Scene {
				return splashscreen.New(logger, assets, screen, input)
			},
		},
	)
	return scenes, splashscreen.SplashScreenSceneId
}
