package main

import (
	"context"
	"embed"
	"fmt"
	"yiya-v2/backend/consts"
	"yiya-v2/backend/services"

	goruntime "runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	isMacOS := goruntime.GOOS == "darwin"

	app := services.AppService()
	fsv := services.FileService()

	width, height := app.GetWindowSize()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            consts.APP_NAME,
		Width:            width,
		Height:           height,
		MinWidth:         consts.MIN_WINDOW_WIDTH,
		MinHeight:        consts.MIN_WINDOW_HEIGHT,
		DisableResize:    true,
		Frameless:        !isMacOS,
		StartHidden:      true,
		AlwaysOnTop:      app.GetIsAlwaysOnTop(),
		BackgroundColour: options.NewRGBA(0, 0, 0, 0),
		DragAndDrop: &options.DragAndDrop{
			EnableFileDrop:     true,
			DisableWebViewDrop: true,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: []interface{}{
			app,
			fsv,
		},
		OnStartup: func(ctx context.Context) {
			app.Setup(ctx)
			fsv.Setup(ctx)

			x, y := app.GetLocalPosition()
			runtime.WindowSetPosition(ctx, x, y)
			runtime.WindowShow(ctx)
		},
		OnDomReady: func(ctx context.Context) {},
		OnBeforeClose: func(ctx context.Context) (prevent bool) {
			app.SaveBoundary()
			return false
		},
		Mac: &mac.Options{
			TitleBar: mac.TitleBarHiddenInset(),
			About: &mac.AboutInfo{
				Title:   fmt.Sprintf("%s v%s", consts.APP_NAME, consts.APP_VERSION),
				Message: "一个小小的图片处理工具.\n\nCopyright © 2024",
				Icon:    icon,
			},
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			DisableZoom:          true,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			DisableFramelessWindowDecorations: false,
			DisablePinchZoom:                  true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
