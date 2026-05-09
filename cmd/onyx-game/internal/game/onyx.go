package game

import (
	"context"

	"github.com/adm87/onyx/pkg/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

type onyx struct {
	ctx    context.Context
	logger *engine.Logger
	screen *engine.Screen
	scenes *engine.Scenes
	time   *engine.Time
}

func newGame(ctx context.Context, cfg *engine.Config, logger *engine.Logger, screen *engine.Screen, scenes *engine.Scenes) *onyx {
	time := engine.NewTime(cfg.FPS)
	return &onyx{
		ctx:    ctx,
		logger: logger,
		screen: screen,
		scenes: scenes,
		time:   time,
	}
}

func (o *onyx) Update() error {
	select {
	case <-o.ctx.Done():
		return o.ctx.Err()
	default:
		o.time.Update()
		return o.scenes.Update(o.time.DeltaTime(), o.time.FixedTime(), o.time.Steps())
	}
}

func (o *onyx) Draw(screen *ebiten.Image) {
	select {
	case <-o.ctx.Done():
		return
	default:
		buffer := o.screen.Buffer()
		buffer.Clear()

		if err := o.scenes.Draw(buffer); err != nil {
			o.logger.Error("error while drawing", "error", err)
		}

		screen.DrawImage(buffer, o.screen.Options())
	}
}

func (o *onyx) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return o.screen.Layout(outsideWidth, outsideHeight)
}
