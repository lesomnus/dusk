package dusk_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/lesomnus/dusk"
	"github.com/stretchr/testify/require"
)

func TestStepHttp(t *testing.T) {
	t.Run("save response to file", func(t *testing.T) {
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Write([]byte("qux"))
		}))
		defer s.Close()

		dst := t.TempDir()
		err := dusk.
			HttpGet(s.URL+"/foo/bar/baz").
			Client(s.Client()).
			Do(t.Context(), nil, dst)
		require.NoError(t, err)

		p := filepath.Join(dst, "baz")
		require.FileExists(t, p)

		data, err := os.ReadFile(p)
		require.NoError(t, err)
		require.Equal(t, "qux", string(data))
	})
}
