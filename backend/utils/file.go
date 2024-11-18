package utils

import (
	"io/fs"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

func IsVaildImage(filePath string) bool {
	info, err := os.Lstat(filePath)
	if err != nil {
		return false
	}

	// 链接文件
	if info.Mode()&os.ModeSymlink != 0 {
		return false
	}

	// 隐藏文件
	if strings.HasPrefix(info.Name(), ".") {
		return false
	}

	ext := filepath.Ext(strings.ToLower(info.Name()))
	m := mime.TypeByExtension(ext)

	supprts := []string{
		"image/jpeg",
		"image/png",
		"image/webp",
	}

	for _, t := range supprts {
		if t == m {
			return true
		}
	}
	return false
}

func GetBaseInfo(path string) types.ImageFileInfo {
	info := types.ImageFileInfo{
		Id:   uuid.NewString(),
		Path: path,
		Size: -1,
	}

	stat, err := os.Stat(path)
	if err == nil {
		info.Size = stat.Size()
		info.Filename = stat.Name()

		ce, err := DetectContentType(path)
		if err == nil {
			info.ContentType = ce
		}
	}

	return info
}

func appendList(path, groupId string, list types.ImageInfoList) types.ImageInfoList {
	vaild := IsVaildImage(path)

	if vaild {
		base := GetBaseInfo(path)
		base.GroupId = groupId
		return append(list, base)
	}
	return list
}

func walk(root, groupId string, list *types.ImageInfoList) {
	_ = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		// 忽略错误
		if err != nil {
			return nil
		}
		// 目录本身不处理
		if path == root {
			return nil
		}

		if info.IsDir() {
			walk(path, groupId, list)
		} else {
			*list = appendList(path, groupId, *list)
		}
		return nil
	})
}

func parsePaths(paths []string, groupId string, list *types.ImageInfoList) {
	for _, path := range paths {
		stat, err := os.Stat(path)
		// 有错就跳过
		if err != nil {
			continue
		}
		if stat.IsDir() {
			walk(path, groupId, list)
		} else {
			*list = appendList(path, groupId, *list)
		}
	}
}

// 解析拖拽进来的路径 输出所有是图片的文件列表 可能存在文件夹 也可能有不支持的文件 都需要过滤
func ParseDropPaths(paths []string) []types.ImageFileInfo {
	groupId := uuid.NewString()
	fileInfos := types.ImageInfoList{}

	parsePaths(paths, groupId, &fileInfos)

	return fileInfos
}
