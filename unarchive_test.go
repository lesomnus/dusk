package dusk_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lesomnus/dusk"
	"github.com/stretchr/testify/require"
)

func TestGunzip(t *testing.T) {
	data, err := testdata.ReadFile("testdata/bar.gz")
	require.NoError(t, err)

	dst := t.TempDir()
	err = os.WriteFile(filepath.Join(dst, "bar.gz"), data, 0644)
	require.NoError(t, err)

	reg := dusk.NewFsRegistry(t.TempDir())
	err = dusk.Gunzip("bar.gz").Do(t.Context(), reg, dst)
	require.NoError(t, err)

	p := filepath.Join(dst, "bar")
	require.FileExists(t, p)

	data, err = os.ReadFile(p)
	require.NoError(t, err)
	require.Equal(t, "foo\n", string(data))
}

func TestUntar(t *testing.T) {
	data, err := testdata.ReadFile("testdata/bar.tar.gz")
	require.NoError(t, err)

	dst := t.TempDir()
	err = os.WriteFile(filepath.Join(dst, "bar.tar.gz"), data, 0644)
	require.NoError(t, err)

	reg := dusk.NewFsRegistry(t.TempDir())
	err = dusk.Untar("bar.tar.gz").Do(t.Context(), reg, dst)
	require.NoError(t, err)

	p := filepath.Join(dst, "bar")
	require.FileExists(t, p)

	data, err = os.ReadFile(p)
	require.NoError(t, err)
	require.Equal(t, "foo\n", string(data))
}
