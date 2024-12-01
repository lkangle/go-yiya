package compress

import (
	"fmt"
	"path/filepath"
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

// DoCompress 执行图片压缩，压缩成功则返回输出的压缩文件路径
func DoCompress(image types.ImageFileInfo, quality int) (string, error) {
	tempdir := utils.GetYiTempDir()
	tempOutput := filepath.Join(tempdir, image.Id)

	lesser := createLessener(image)
	if lesser == nil {
		return "", fmt.Errorf("不支持的图片类型: %s", image.ContentType)
	}

	err := lesser.Compress(image.Path, tempOutput, quality)
	if err != nil {
		return "", err
	}

	return tempOutput, nil
}
