package dusk

import (
	"context"
)

type Package interface {
	Name() string
	Version() string

	// Deps() []Package
	Unpack(ctx context.Context, reg Registry, dst string) ([]Exec, error)
}
