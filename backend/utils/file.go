package utils

import (
	"net/http"
	"os"
	"yiya-v2/backend/types"

	"github.com/google/uuid"
)

func DetectContentType(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	// 读取前 512 个字节
	buffer := make([]byte, 512)
	_, err = f.Read(buffer)
	if err != nil {
		return "", err
	}

	// 使用 http 包检测 MIME 类型
	return http.DetectContentType(buffer), nil
}

func GetBaseInfo(path string) types.ImageFileInfo {
	info := types.ImageFileInfo{
		Id:         uuid.NewString(),
		Path:       path,
		OriginSize: -1,
		Size:       0,
	}

	stat, err := os.Stat(path)
	if err == nil {
		info.OriginSize = stat.Size()
		info.Filename = stat.Name()

		ce, err := DetectContentType(path)
		if err == nil {
			info.ContentType = ce
		}
	}

	return info
}
