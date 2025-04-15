package dusk

import (
	"context"
	"os"
	"path/filepath"

	"github.com/lesomnus/z"
)

type FsReg struct {
	root  string
	repos map[string]Repo
}

func NewFsRegistry(root string) *FsReg {
	return &FsReg{
		root:  root,
		repos: map[string]Repo{},
	}
}

func (r *FsReg) Repo(name string) (Repo, bool) {
	m, ok := r.repos[name]
	return m, ok
}

func (r *FsReg) path(pkg Package) string {
	return filepath.Join(r.root, pkg.Name(), pkg.Version())
}

func (r *FsReg) register(v Exec) {
	m, ok := r.repos[v.Name()]
	if !ok {
		m = map[string]Exec{}
		r.repos[v.Name()] = m
	}

	m[v.Version()] = v
}

func (r *FsReg) Register(execs ...Exec) {
	for _, v := range execs {
		r.register(v)
	}
}

func (r *FsReg) Resolve(ctx context.Context, pkg Package) ([]Exec, error) {
	p := r.path(pkg)
	if err := os.MkdirAll(p, 0755); err != nil {
		return nil, z.WrapErr("make dir", err)
	}

	vs, err := pkg.Unpack(ctx, r, p)
	if err != nil {
		return nil, z.WrapErr("unpack", err)
	}

	r.Register(vs...)
	return vs, nil
}
