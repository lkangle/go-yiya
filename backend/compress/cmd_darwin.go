//go:build darwin
// +build darwin

package compress

import (
	"context"
	"os/exec"
)

func createCommandContext(ctx context.Context, name string, arg ...string) *exec.Cmd {
	return exec.CommandContext(ctx, name, arg...)
}
