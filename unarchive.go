package dusk

import (
	"context"
	"path/filepath"

	"github.com/lesomnus/z"
)

func Gunzip(rpath string) Step {
	return StepFunc(func(ctx context.Context, reg Registry, dst string) error {
		gzip_pkg := LookPath("gzip")
		gzip, err := Resolve(ctx, reg, gzip_pkg)
		if err != nil {
			return z.WrapErr("resolve gzip", err)
		}

		src_path := filepath.Join(dst, rpath)
		cmd := gzip.Cmd(ctx, "--decompress", src_path)
		return cmd.Run()
	})
}

func Untar(rpath string) Step {
	return StepFunc(func(ctx context.Context, reg Registry, dst string) error {
		tar_pkg := LookPath("tar")
		tar, err := Resolve(ctx, reg, tar_pkg)
		if err != nil {
			return z.WrapErr("resolve tar", err)
		}

		src_path := filepath.Join(dst, rpath)
		cmd := tar.Cmd(ctx, "--directory", dst, "-xf", src_path)
		return cmd.Run()
	})
}
