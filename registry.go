package dusk

import (
	"context"
	"iter"
	"maps"
)

type Registry interface {
	Repo(name string) (Repo, bool)
	Register(execs ...Exec)
	Resolve(ctx context.Context, pkg Package) ([]Exec, error)
}

type Repo map[string]Exec // key is version.

func (r Repo) Get(version string) (Exec, bool) {
	v, ok := r[version]
	return v, ok
}

func Resolve(ctx context.Context, reg Registry, pkg Package) (Exec, error) {
	// TODO: from option
	name := pkg.Name()

	if repo, ok := reg.Repo(name); ok && len(repo) > 0 {
		next, stop := iter.Pull(maps.Values(repo))
		defer stop()

		v, _ := next()
		return v, nil
	}

	vs, err := reg.Resolve(ctx, pkg)
	if err != nil {
		return nil, err
	}

	for _, v := range vs {
		if v.Name() == name {
			return v, nil
		}
	}

	return nil, nil
}
