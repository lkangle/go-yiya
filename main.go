package main

import (
	"context"
	"embed"
	"yiya-v2/backend/consts"
	"yiya-v2/backend/services"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// isMacOS := goruntime.GOOS == "darwin"

	bdy := services.Boundary()
	width, height := bdy.GetWindowSize()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "易压v2",
		Width:     width,
		Height:    height,
		MinWidth:  consts.MIN_WINDOW_WIDTH,
		MinHeight: consts.MIN_WINDOW_HEIGHT,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        func(ctx context.Context) {},
		Bind:             []interface{}{},
		OnBeforeClose: func(ctx context.Context) (prevent bool) {
			bdy.UpdateBoundary(ctx)
			return false
		},
		OnDomReady: func(ctx context.Context) {
			x, y := bdy.GetWindowPosition(ctx)
			runtime.WindowSetPosition(ctx, x, y)
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
