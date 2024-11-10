package services

import (
	"sync"
	"yiya-v2/backend/types"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/net/context"
)

type fileService struct {
	ctx context.Context
}

var fsv *fileService
var fsonce sync.Once

func FileService() *fileService {
	if fsv == nil {
		fsonce.Do(func() {
			fsv = &fileService{
				ctx: nil,
			}
		})
	}
	return fsv
}

func (f *fileService) Setup(ctx context.Context) {
	f.ctx = ctx
}

func (fs *fileService) OpenSelectFilesDialog() (resp types.JSResp) {
	fileList, err := runtime.OpenMultipleFilesDialog(fs.ctx, runtime.OpenDialogOptions{
		DefaultFilename:            "",
		Title:                      "选择图片进行压缩",
		ShowHiddenFiles:            false,
		CanCreateDirectories:       false,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: true,
		Filters: []runtime.FileFilter{
			{
				Pattern: "*.png;*.jpg;*.jpeg;*.webp",
			},
		},
	})

	if err != nil {
		resp.Success = false
		resp.Msg = err.Error()
		return
	}

	resp.Data = fileList
	resp.Success = true
	return
}

func (fs *fileService) CompressImages(paths []string, opt types.CompressOptions) (resp types.JSResp) {

	resp.Success = true
	return
}
