package cli

import (
	"flag"

	"github.com/adm87/onyx/internal/game/scenes/splashscreen"
)

type GameArgs struct {
	Fullscreen bool
	Scene      string
}

func NewGameArgs() *GameArgs {
	return &GameArgs{
		Fullscreen: false,
		Scene:      string(splashscreen.SceneId),
	}
}

func (ga *GameArgs) Parse(prog string, args []string) error {
	set := flag.NewFlagSet(prog, flag.ContinueOnError)
	set.BoolVar(&ga.Fullscreen, "fullscreen", ga.Fullscreen, "run the game in fullscreen mode")
	set.StringVar(&ga.Scene, "scene", ga.Scene, "the initial scene to start the game with")

	if err := set.Parse(args); err != nil {
		return err
	}
	return nil
}
