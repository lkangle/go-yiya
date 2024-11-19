package compress

import (
	"fmt"
	"path"
	"yiya-v2/backend/types"
	"yiya-v2/backend/utils"
)

type lessener interface {
	Compress(input, output string, quality int) error
}

var pool = make(map[string](lessener), 3)

func createLessener(image types.ImageFileInfo) lessener {
	ctype := image.ContentType
	er, ok := pool[ctype]
	if ok {
		return er
	}

	if ctype == "image/jpeg" {
		j := &jpegLessener{}
		pool[ctype] = j
		return j
	}

	if ctype == "image/png" {
		p := &pngLessener{}
		pool[ctype] = p
		return p
	}

	if ctype == "image/webp" {
		w := &webpLessener{}
		pool[ctype] = w
		return w
	}
	return nil
}

func DoCompress(image types.ImageFileInfo, option types.CompressOptions) (string, error) {
	tempdir := utils.GetImageTempDir(image.Id)

	tempOutput := path.Join(tempdir, "compress_image.data")

	lesser := createLessener(image)
	if lesser == nil {
		return "", fmt.Errorf("不支持的图片类型: %s", image.ContentType)
	}

	err := lesser.Compress(image.Path, tempOutput, option.Quality)
	if err != nil {
		return "", err
	}

	// TODO: 各种文件处理操作
	return tempOutput, nil
}
