package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"yiya-v2/backend/consts"
	"yiya-v2/backend/services"
	"yiya-v2/backend/utils"

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

	opts := utils.GetOptionUtils()
	sys := services.System()

	width, height := opts.GetWindowSize()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         consts.APP_NAME,
		Width:         width,
		Height:        height,
		MinWidth:      consts.MIN_WINDOW_WIDTH,
		MinHeight:     consts.MIN_WINDOW_HEIGHT,
		DisableResize: true,
		Frameless:     !isMacOS,
		StartHidden:   true,
		AlwaysOnTop:   opts.IsAlwaysOnTop(),
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind: []interface{}{
			sys,
		},
		OnStartup: func(ctx context.Context) {
			x, y := opts.GetWindowPosition(ctx)
			runtime.WindowSetPosition(ctx, x, y)
			runtime.WindowShow(ctx)
			sys.Setup(ctx)
		},
		OnDomReady: func(ctx context.Context) {},
		OnBeforeClose: func(ctx context.Context) (prevent bool) {
			opts.UpdateBoundary(ctx)

			log.Fatalln("==> ...这里还关闭了...")
			return false
		},
		Mac: &mac.Options{
			TitleBar: mac.TitleBarHiddenInset(),
			About: &mac.AboutInfo{
				Title:   fmt.Sprintf("%s %s", consts.APP_NAME, consts.APP_VERSION),
				Message: "一个小小的图片处理工具.\n\nCopyright © 2024",
				Icon:    icon,
			},
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              false,
			WindowIsTranslucent:               false,
			DisableFramelessWindowDecorations: false,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
