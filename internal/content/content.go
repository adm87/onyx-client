package content

import (
	"embed"

	"github.com/adm87/onyx/pkg/engine"
)

//go:embed static
var StaticAssetFS embed.FS

const (
	StaticSplashScreenFilePath = engine.FilePath("static/splash_1920x1080_black.png")
	StatisImg10x10FilePath     = engine.FilePath("static/img_10x10.png")
)
