package utils

import (
	"path/filepath"
	"runtime"

	"github.com/vrischmann/userdir"
)

var IsMacOS bool

func GetYiHome() string {
	return filepath.Join(userdir.GetConfigHome(), "YIYA_IG")
}

func GetYiTempDir() string {
	return filepath.Join(GetYiHome(), "temp")
}

func init() {
	IsMacOS = runtime.GOOS == "darwin"
}
