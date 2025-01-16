package main

import (
	"embed"
	"github.com/run-bigpig/conduit/internal/bootstrap"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	boot := bootstrap.NewBoot()
	err := wails.Run(&options.App{
		Title:     "conduit",
		Width:     1024,
		Height:    768,
		MinHeight: 768,
		MinWidth:  1024,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        boot.Startup,
		Bind: []interface{}{
			boot.NewGuiHandler(),
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
