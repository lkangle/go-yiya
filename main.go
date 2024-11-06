package main

import (
	"context"
	"embed"
	"fmt"
	"yiya-v2/backend/consts"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "易压v2",
		Width:     consts.DEFAULT_WINDOW_WIDTH,
		Height:    consts.DEFAULT_WINDOW_HEIGHT,
		MinWidth:  consts.MIN_WINDOW_WIDTH,
		MinHeight: consts.MIN_WINDOW_HEIGHT,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		OnBeforeClose: func(ctx context.Context) (prevent bool) {
			s, _ := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
				Title:   "真的要关闭吗？",
				Message: "关闭应用就打不开了哈",
				Buttons: []string{
					"确定",
				},
			})

			fmt.Println("xxxxxsss", s)
			return false
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
