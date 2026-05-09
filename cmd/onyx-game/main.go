package main

import (
	"context"
	"errors"
	"os"

	"github.com/adm87/onyx/cmd/onyx-game/internal/game"
	"github.com/adm87/onyx/cmd/onyx-game/internal/game/cli"
	"github.com/adm87/onyx/pkg/engine"
)

var version = "0.0.0-unreleased"

func main() {
	args := cli.NewGameArgs()
	if err := args.Parse(os.Args[0], os.Args[1:]); err != nil {
		println("error parsing arguments:", err.Error())
		os.Exit(1)
	}
	cfg := &engine.Config{
		Title:        "Onyx Game",
		Width:        1280,
		Height:       720,
		FPS:          1,
		Fullscreen:   args.Fullscreen,
		InitialScene: args.Scene,
	}
	if err := game.Boot(cfg); err != nil && !errors.Is(err, context.Canceled) {
		println("error running game:", err.Error())
		os.Exit(1)
	}
}
