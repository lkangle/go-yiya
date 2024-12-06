//go:build windows
// +build windows

package compress

import (
	"context"
	"os/exec"
	"syscall"
)

func createCommandContext(ctx context.Context, name string, arg ...string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, name, arg...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	return cmd
}
