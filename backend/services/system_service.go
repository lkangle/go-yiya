package services

import (
	"sync"
	"yiya-v2/backend/consts"
	"yiya-v2/backend/types"
	"yiya-v2/backend/utils"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/net/context"
)

type systimeService struct {
	ctx context.Context
}

var system *systimeService
var once sync.Once

func System() *systimeService {
	if system == nil {
		once.Do(func() {
			system = &systimeService{
				ctx: nil,
			}
		})
	}
	return system
}

func (s *systimeService) Setup(ctx context.Context) {
	s.ctx = ctx
}

func (s *systimeService) GetAppInfo() types.JsObject {
	return map[string]any{
		"name":    consts.APP_NAME,
		"varsion": consts.APP_VERSION,
	}
}

// 设置是否置顶窗口
func (s *systimeService) SetIsAllowsOnTop(top bool) {
	runtime.WindowSetAlwaysOnTop(s.ctx, top)
	utils.GetOptionUtils().SaveIsAlwaysOnTop(top)
}

// 获取当前是否置顶窗口
func (s *systimeService) GetIsAllowsOnTop() bool {
	return utils.GetOptionUtils().IsAlwaysOnTop()
}
