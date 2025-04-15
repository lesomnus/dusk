package dusk

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/lesomnus/z"
)

type HttpStep struct {
	req    *http.Request
	client *http.Client
}

func Http(req *http.Request) *HttpStep {
	return &HttpStep{req, http.DefaultClient}
}

func (s *HttpStep) Client(c *http.Client) *HttpStep {
	s.client = c
	return s
}

func (s *HttpStep) Do(ctx context.Context, reg Registry, dst string) error {
	res, err := s.client.Do(s.req.WithContext(ctx))
	if err != nil {
		return z.WrapErr("HTTP GET", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", res.Status)
	}

	out_name := path.Base(s.req.URL.Path)
	if out_name == "" || out_name == "/" || out_name == "." {
		out_name = "body"
	}

	out_path := filepath.Join(dst, out_name)
	out_file, err := os.Create(out_path)
	if err != nil {
		return z.WrapErr("create", err)
	}
	defer out_file.Close()

	if _, err := io.Copy(out_file, res.Body); err != nil {
		return z.WrapErr("copy body", err)
	}

	return nil
}

func HttpGet(url string) *HttpStep {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(z.WrapErr("new request", err))
	}

	return Http(req)
}
