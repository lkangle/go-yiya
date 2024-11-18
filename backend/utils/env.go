package utils

import (
	"runtime"
)

var IsMacOS bool

func init() {
	IsMacOS = runtime.GOOS == "darwin"
}
