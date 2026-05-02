package bindings

import (
	"github.com/adm87/onyx/pkg/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Quit engine.InputBindingID = iota
	Fullscreen
)

var keyboardBindingTable = map[engine.InputBindingID]ebiten.Key{
	Quit:       ebiten.KeyEscape,
	Fullscreen: ebiten.KeyF11,
}
