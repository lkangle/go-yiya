package utils

import (
	"fmt"
	"io"
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

func IsValidImage(filePath string) bool {
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

func GetSize(path string) int64 {
	stat, err := os.Stat(path)
	if err == nil {
		return stat.Size()
	}
	return 0
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

func appendList(path, rootDir string, list []types.ImageFileInfo) []types.ImageFileInfo {
	valid := IsValidImage(path)

	if valid {
		base := GetBaseInfo(path)
		base.DirPath = rootDir
		return append(list, base)
	}
	return list
}

func walk(root, rootDir string, list *[]types.ImageFileInfo) {
	_ = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		// 忽略错误
		if err != nil {
			return nil
		}
		// 目录本身不处理
		if path == root {
			return nil
		}

		if d.IsDir() {
			walk(path, rootDir, list)
		} else {
			*list = appendList(path, rootDir, *list)
		}
		return nil
	})
}

// ParseDropPaths 解析拖拽进来的路径 输出所有是图片的文件列表 可能存在文件夹 也可能有不支持的文件 都需要过滤
func ParseDropPaths(paths []string) []types.ImageFileInfo {
	list := new([]types.ImageFileInfo)

	for _, path := range paths {
		info, err := os.Stat(path)
		// 有错就跳过
		if err != nil {
			continue
		}
		if info.IsDir() {
			walk(path, path, list)
		} else {
			*list = appendList(path, "", *list)
		}
	}

	return *list
}

func CreateFile(path string) (file *os.File, err error) {
	dir := filepath.Dir(path)

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	file, err = os.Create(path)
	return
}

// MoveFile 使用 io.Copy 来实现文件复制并删除源文件
func MoveFile(dst, src string) error {
	err := CopyFile(dst, src)
	if err != nil {
		return err
	}

	// 删除源文件
	err = os.Remove(src)
	if err != nil {
		return fmt.Errorf("failed to remove source file: %w", err)
	}
	return nil
}

func CopyFile(dst, src string) error {
	ipt, err := os.Open(src)
	if err != nil {
		return err
	}
	defer ipt.Close()

	out, err := CreateFile(dst)
	if err != nil {
		return err
	}
	defer func() {
		out.Close()
		if err != nil {
			os.Remove(dst)
		}
	}()

	_, err = io.Copy(out, ipt)
	return err
}

// CopyToTemp 复制文件到临时目录下，并返回临时文件路径
func CopyToTemp(input string) string {
	f := filepath.Join(GetYiTempDir(), uuid.NewString())
	err := CopyFile(f, input)
	if err != nil {
		return ""
	}
	return f
}

// getNewPathWithSuffix 根据输入路径和后缀返回新的路径
func getNewPathWithSuffix(input, suffix string) string {
	// 获取路径的基本文件名和扩展名
	base := filepath.Base(input)
	ext := filepath.Ext(base)

	// 如果路径是文件，则在文件名的扩展名之前添加后缀
	if ext != "" {
		// 对文件路径进行处理：添加后缀
		dir := filepath.Dir(input)
		newBase := strings.TrimSuffix(base, ext) + "-" + suffix + ext
		return filepath.Join(dir, newBase)
	}

	// 如果是目录，则直接在目录名后添加后缀
	return input + "-" + suffix
}

func GetOutputWithSuffix(info types.ImageFileInfo, suffix string) string {
	if suffix == "" {
		suffix = "yiya"
	}

	if info.DirPath != "" {
		dir := info.DirPath
		relPath, err := filepath.Rel(dir, info.Path)
		if err != nil {
			return ""
		}
		suffixDir := getNewPathWithSuffix(dir, suffix)
		return filepath.Join(suffixDir, relPath)
	}

	return getNewPathWithSuffix(info.Path, suffix)
}
