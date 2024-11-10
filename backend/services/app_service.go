package services

import (
	"sync"
	"yiya-v2/backend/consts"
	"yiya-v2/backend/storage"
	"yiya-v2/backend/types"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/net/context"
)

type appService struct {
	ctx context.Context
	opt *storage.OptionsStorage
}

var apps *appService
var appOnce sync.Once

func AppService() *appService {
	if apps == nil {
		appOnce.Do(func() {
			apps = &appService{
				ctx: nil,
				opt: storage.NewOptionStorage(),
			}
		})
	}
	return apps
}

func (s *appService) Setup(ctx context.Context) {
	s.ctx = ctx
}

func (s *appService) GetAppInfo() types.JsObject {
	return map[string]any{
		"name":    consts.APP_NAME,
		"varsion": consts.APP_VERSION,
	}
}

func (o *appService) GetIsAlwaysOnTop() bool {
	return o.opt.Get().System.IsAlwaysOnTop
}

func (p *appService) SetIsAlwaysOnTop(top bool) {
	old := p.opt.Get()

	runtime.WindowSetAlwaysOnTop(p.ctx, top)

	old.System.IsAlwaysOnTop = top

	p.opt.Save(&old)
}

func (p *appService) GetWindowSize() (width, height int) {
	data := p.opt.Get().System
	width, height = data.WindowWidth, data.WindowHeight
	if width <= 0 {
		width = consts.DEFAULT_WINDOW_WIDTH
	}
	if height <= 0 {
		height = consts.DEFAULT_WINDOW_HEIGHT
	}
	return
}

func (p *appService) GetLocalPosition() (x, y int) {
	data := p.opt.Get().System
	x, y = data.WindowPosX, data.WindowPosY
	width, height := data.WindowWidth, data.WindowHeight
	var screenWidth, screenHeight int
	if screens, err := runtime.ScreenGetAll(p.ctx); err == nil {
		for _, screen := range screens {
			if screen.IsCurrent {
				screenWidth, screenHeight = screen.Size.Width, screen.Size.Height
				break
			}
		}
	}
	if screenWidth <= 0 || screenHeight <= 0 {
		screenWidth, screenHeight = consts.DEFAULT_WINDOW_WIDTH, consts.DEFAULT_WINDOW_HEIGHT
	}
	if x <= 0 || x+width > screenWidth || y <= 0 || y+height > screenHeight {
		// out of screen, reset to center
		x, y = (screenWidth-width)/2, (screenHeight-height)/2
	}
	return
}

func (p *appService) SaveBoundary() {
	old := p.opt.Get()

	sys := old.System

	width, height := runtime.WindowGetSize(p.ctx)
	x, y := runtime.WindowGetPosition(p.ctx)

	if width >= consts.MIN_WINDOW_WIDTH && height >= consts.MIN_WINDOW_HEIGHT {
		sys.WindowWidth = width
		sys.WindowHeight = height
	}
	sys.WindowPosX = x
	sys.WindowPosY = y
	old.System = sys

	p.opt.Save(&old)
}

func (p *appService) GetCompressOptions() types.JSResp {
	opt := p.opt.Get()

	return types.JSResp{
		Success: true,
		Data:    opt.Compress,
	}
}

func (p *appService) UpdateCompressOptions(op types.CompressOptions) (resp types.JSResp) {
	old := p.opt.Get()
	oc := old.Compress

	oc.Override = op.Override
	if op.Quality == consts.BEST_QUALITY || op.Quality == consts.NORMAL_QUALITY {
		oc.Quality = op.Quality
	}
	if op.NewSuffix != "" {
		oc.NewSuffix = op.NewSuffix
	}

	old.Compress = oc

	resp.Success = true
	return
}
