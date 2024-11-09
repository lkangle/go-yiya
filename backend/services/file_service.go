package services

import (
	"sync"

	"golang.org/x/net/context"
)

type fileService struct {
	ctx context.Context
}

var fsv *fileService
var fsonce sync.Once

func NewFileService() *fileService {
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

func (fs *fileService) OpenSelectFilesDialog() {

}
