package services

import (
	"context"
	"yiya-v2/backend/consts"
	"yiya-v2/backend/storage"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type boundaryService struct {
	bs *storage.BoundaryStorage
}

func Boundary() *boundaryService {
	return &boundaryService{
		bs: storage.NewBoundaryStorage(),
	}
}

func (p *boundaryService) GetWindowSize() (width, height int) {
	data := p.bs.Get()
	width, height = data.WindowWidth, data.WindowHeight
	if width <= 0 {
		width = consts.DEFAULT_WINDOW_WIDTH
	}
	if height <= 0 {
		height = consts.DEFAULT_WINDOW_HEIGHT
	}
	return
}

func (p *boundaryService) GetWindowPosition(ctx context.Context) (x, y int) {
	data := p.bs.Get()
	x, y = data.WindowPosX, data.WindowPosY
	width, height := data.WindowWidth, data.WindowHeight
	var screenWidth, screenHeight int
	if screens, err := runtime.ScreenGetAll(ctx); err == nil {
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

func (p *boundaryService) UpdateBoundary(ctx context.Context) {
	old := p.bs.Get()

	width, height := runtime.WindowGetSize(ctx)
	x, y := runtime.WindowGetPosition(ctx)

	if width >= consts.MIN_WINDOW_WIDTH && height >= consts.MIN_WINDOW_HEIGHT {
		old.WindowWidth = width
		old.WindowHeight = height
	}
	old.WindowPosX = x
	old.WindowPosY = y

	p.bs.Save(&old)
}
