package dusk

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/lesomnus/z"
)

type OsFile struct {
	descriptor
	path string
}

func Os(path string) *OsFile {
	return &OsFile{
		descriptor: descriptor{
			name:    filepath.Base(path),
			version: "os",
		},
		path: path,
	}
}

func Static(name, version string, path string) *OsFile {
	return &OsFile{
		descriptor: descriptor{
			name:    name,
			version: version,
		},
		path: path,
	}
}

func LookPath(name string) *OsFile {
	p, err := exec.LookPath(name)
	if err != nil {
		p = "/not-exists"
	}

	return Os(p)
}

func (p *OsFile) Unpack(ctx context.Context, reg Registry, dst string) ([]Exec, error) {
	info, err := os.Stat(p.path)
	if err != nil {
		return nil, z.WrapErr("stat", err)
	}

	_ = info

	// TODO: check if the file can be executed by current user.

	return []Exec{p}, nil
}

func (p *OsFile) Cmd(ctx context.Context, args ...string) *exec.Cmd {
	return exec.CommandContext(ctx, p.path, args...)
}
