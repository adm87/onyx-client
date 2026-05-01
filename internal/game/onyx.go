package game

import (
	"context"

	"github.com/adm87/onyx/pkg/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

type onyx struct {
	ctx    context.Context
	logger *engine.Logger
	assets *engine.Assets
	input  *engine.Input
	screen *engine.Screen
	scenes *engine.Scenes
}

func newGame(
	ctx context.Context,
	cfg *engine.Config,
	logger *engine.Logger,
	input *engine.Input,
	assets *engine.Assets,
	screen *engine.Screen,
	scenes *engine.Scenes) *onyx {
	return &onyx{
		ctx:    ctx,
		logger: logger,
		assets: assets,
		input:  input,
		screen: screen,
		scenes: scenes,
	}
}

func (o *onyx) Update() error {
	select {
	case <-o.ctx.Done():
		return o.ctx.Err()
	default:
		if err := o.input.Poll(); err != nil {
			return err
		}
		if err := o.scenes.Update(); err != nil {
			return err
		}
		return nil
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
			o.logger.Error("error while drawing: %v", err)
		}

		screen.DrawImage(buffer, o.screen.Options())
	}
}

func (o *onyx) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return o.screen.Layout(outsideWidth, outsideHeight)
}
