package utils

import (
	"path"
	"runtime"

	"github.com/vrischmann/userdir"
)

var IsMacOS bool

func GetYiHome() string {
	return path.Join(userdir.GetConfigHome(), "YIYA_IG")
}

func GetImageTempDir(id string) string {
	return path.Join(GetYiHome(), "compress_temp", id)
}

func init() {
	IsMacOS = runtime.GOOS == "darwin"
}
