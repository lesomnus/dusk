package dusk

import (
	"context"
	"os"
	"os/exec"
)

type Exec interface {
	Name() string
	Version() string

	Cmd(ctx context.Context, args ...string) *exec.Cmd
}

func Run(ctx context.Context, e Exec, args ...string) error {
	cmd := e.Cmd(ctx, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
