package utils

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func GetPngquantFilename() string {
	goos := runtime.GOOS
	arch := runtime.GOARCH

	switch goos {
	case "windows":
		return "pngquant-windows.exe"
	case "darwin":
		return "pngquant-darwin-" + arch
	case "linux":
		return "pngquant-linux"
	}
	return ""
}

// GetPngquantFullPath 获取可执行文件完整路径
func GetPngquantFullPath() string {
	home := GetYiHome()
	pngquantPath := filepath.Join(home, GetPngquantFilename())

	return pngquantPath
}

// CompareVersions 比较两个版本号 a 和 b
// 返回值:
// -1: a < b
//
//	0: a == b
//	1: a > b
func CompareVersions(a, b string) int {
	aParts := strings.Split(a, ".")
	bParts := strings.Split(b, ".")

	// 找出较长的版本号长度
	maxLen := len(aParts)
	if len(bParts) > maxLen {
		maxLen = len(bParts)
	}

	for i := 0; i < maxLen; i++ {
		var aNum, bNum int

		// 转换每个部分的数字，超出长度默认为 0
		if i < len(aParts) {
			aNum, _ = strconv.Atoi(aParts[i])
		}
		if i < len(bParts) {
			bNum, _ = strconv.Atoi(bParts[i])
		}

		// 比较当前部分的数字
		if aNum < bNum {
			return -1
		}
		if aNum > bNum {
			return 1
		}
	}

	// 如果循环结束后没有发现差异，版本号相等
	return 0
}

func DownloadFile(url, filepath string) error {
	// 创建 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("无法发送 GET 请求: %w", err)
	}
	defer resp.Body.Close()

	// 检查 HTTP 响应状态
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载失败，HTTP 状态码: %d", resp.StatusCode)
	}

	// 创建目标文件
	out, err := CreateFile(filepath)
	if err != nil {
		return fmt.Errorf("无法创建文件: %w", err)
	}
	defer out.Close()

	// 将响应数据写入文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("保存文件失败: %w", err)
	}

	return nil
}
