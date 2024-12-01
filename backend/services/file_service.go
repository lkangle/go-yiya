package services

import (
	"os/exec"
	"sync"
	"yiya-v2/backend/compress"
	"yiya-v2/backend/types"
	"yiya-v2/backend/utils"

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

func (f *fileService) ParseFilePaths(paths []string) []types.ImageFileInfo {
	list := utils.ParseDropPaths(paths)
	return list
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

	fileInfos := utils.Map(fileList, func(path string) types.ImageFileInfo {
		return utils.GetBaseInfo(path)
	})
	resp.Data = fileInfos
	resp.Success = true
	return
}

func (fs *fileService) OpenFile(path string) (resp types.JSResp) {
	var cmd *exec.Cmd
	if utils.IsMacOS {
		cmd = exec.Command("open", "-R", path)
	} else {
		cmd = exec.Command("explorer", "/select,", path)
	}

	err := cmd.Run()
	resp.Success = true
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); !ok || exitErr.ExitCode() != 1 {
			resp.Success = false
			resp.Msg = err.Error()
		}
	}
	return
}

func (fs *fileService) CompressImage(input types.ImageFileInfo, opt types.CompressOptions) (res types.CompressResult) {
	outPath, err := compress.DoCompress(input, opt.Quality)

	if err != nil {
		res.Message = err.Error()
		res.Success = false
	} else {
		res.Success = true
		res.Path = outPath
		size := utils.GetSize(outPath)
		res.Size = int(size)
	}
	return
}

func (fs *fileService) RestoreImage(input types.ImageFileInfo) (resp types.JSResp) {
	return
}
