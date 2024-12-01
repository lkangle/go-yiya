package compress

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"
	"time"
	"yiya-v2/backend/types"
	"yiya-v2/backend/utils"
)

type pngLessener struct {
}

func (j *pngLessener) Compress(input, output string, quality int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	outdir := filepath.Dir(output)
	_ = os.MkdirAll(outdir, os.ModePerm)

	maxQua := j.mapQuality(quality)

	binpath := utils.GetPngquantBinPath()
	cmd := exec.CommandContext(
		ctx,
		binpath,
		"-f",
		"--speed=3",
		"--quality=10-"+maxQua,
		"--floyd=1",
		"-o", output,
		input,
	)

	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow: true,
		}
	}

	out, err := cmd.CombinedOutput()
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return ctx.Err()
	}
	outstr := string(out)
	if outstr != "" {
		return errors.New(outstr)
	}
	return err
}

func (j *pngLessener) mapQuality(q int) string {
	if q == types.QuaNormal {
		return "90"
	}
	return "98"
}
