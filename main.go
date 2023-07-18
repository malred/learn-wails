package main

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"rp/include"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	cal := include.NewCalculator()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "rp",
		Width:  1024,
		Height: 768, 
		// 菜单(没搞懂)
		// Menu: app.menu(),
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		// OnStartup:        app.startup,
		OnStartup: func(ctx context.Context) {
			app.SetContext(ctx)
			cal.SetContext(ctx)
		},
		Bind: []interface{}{
			app,
			cal,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
