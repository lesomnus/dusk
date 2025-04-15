package dusk_test

import (
	"context"
	"testing"

	"github.com/lesomnus/dusk"
	"github.com/stretchr/testify/require"
)

func TestPipe(t *testing.T) {
	vs := []string{}
	step := func(v string) dusk.Step {
		return dusk.StepFunc(func(ctx context.Context, reg dusk.Registry, d string) error {
			vs = append(vs, v)
			return nil
		})
	}

	pkg := dusk.Pipe("foo", "bar",
		step("foo"),
		step("bar"),
		step("baz"),
	).Sink(func(des dusk.Descriptor, dst string) ([]dusk.Exec, error) {
		return []dusk.Exec{dusk.Os("/bin/cat")}, nil
	})
	require.Equal(t, "foo", pkg.Name())
	require.Equal(t, "bar", pkg.Version())

	execs, err := pkg.Unpack(t.Context(), nil, "")
	require.NoError(t, err)
	require.Equal(t, []string{"foo", "bar", "baz"}, vs)
	require.Len(t, execs, 1)

	exec := execs[0]
	require.Equal(t, "cat", exec.Name())
	require.Equal(t, "os", exec.Version())

	cmd := exec.Cmd(t.Context())
	require.Equal(t, "/bin/cat", cmd.Path)
}
