package engine

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type RendererType string

type Renderable interface {
	RendererType() RendererType
}

type RendererAdapter interface {
	Type() RendererType
	Render(screen *ebiten.Image, renderable Renderable, matrix ebiten.GeoM) error
}

type Renderer struct {
	logger   *Logger
	adapters map[RendererType]RendererAdapter
}

func NewRenderer(logger *Logger) *Renderer {
	return &Renderer{
		logger:   logger,
		adapters: make(map[RendererType]RendererAdapter),
	}
}

func (r *Renderer) AddAdapter(adapter RendererAdapter) {
	r.adapters[adapter.Type()] = adapter
}

func (r *Renderer) Render(screen *ebiten.Image, renderable Renderable, matrix ebiten.GeoM) error {
	if adapter, exists := r.adapters[renderable.RendererType()]; exists {
		return adapter.Render(screen, renderable, matrix)
	}
	return fmt.Errorf("no adapter found for renderer type: %s", renderable.RendererType())
}
