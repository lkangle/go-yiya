package utils

import (
	"context"
	"sync"
	"yiya-v2/backend/consts"
	"yiya-v2/backend/storage"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type optionUtils struct {
	opts *storage.OptionStorage
}

var util *optionUtils
var oncev sync.Once

func GetOptionUtils() *optionUtils {
	if util == nil {
		oncev.Do(func() {
			util = &optionUtils{
				opts: storage.NewOptionStorage(),
			}
		})
	}

	return util
}

func (p *optionUtils) IsAlwaysOnTop() bool {
	return p.opts.Get().IsAllowsOnTop
}

func (p *optionUtils) SaveIsAlwaysOnTop(top bool) {
	old := p.opts.Get()

	old.IsAllowsOnTop = top

	p.opts.Save(&old)
}

func (p *optionUtils) GetWindowSize() (width, height int) {
	data := p.opts.Get()
	width, height = data.WindowWidth, data.WindowHeight
	if width <= 0 {
		width = consts.DEFAULT_WINDOW_WIDTH
	}
	if height <= 0 {
		height = consts.DEFAULT_WINDOW_HEIGHT
	}
	return
}

func (p *optionUtils) GetWindowPosition(ctx context.Context) (x, y int) {
	data := p.opts.Get()
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

func (p *optionUtils) UpdateBoundary(ctx context.Context) {
	old := p.opts.Get()

	width, height := runtime.WindowGetSize(ctx)
	x, y := runtime.WindowGetPosition(ctx)

	if width >= consts.MIN_WINDOW_WIDTH && height >= consts.MIN_WINDOW_HEIGHT {
		old.WindowWidth = width
		old.WindowHeight = height
	}
	old.WindowPosX = x
	old.WindowPosY = y

	p.opts.Save(&old)
}
