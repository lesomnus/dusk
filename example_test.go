package dusk_test

import (
	"context"
	"os"
	"time"

	"github.com/lesomnus/dusk"
)

func Example_resolveRemoteCli() {
	reg := dusk.NewFsRegistry(os.TempDir() + "/dusk_example")
	pkg := dusk.Pipe("oras", "v1.2.2",
		dusk.HttpGet("https://github.com/oras-project/oras/releases/download/v1.2.2/oras_1.2.2_linux_amd64.tar.gz"),
		dusk.Untar("oras_1.2.2_linux_amd64.tar.gz")).
		Pick("oras")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exec, err := dusk.Resolve(ctx, reg, pkg)
	if err != nil {
		panic(err)
	}

	cmd := exec.Cmd(ctx, "version")
	cmd.Stdout = os.Stdout
	cmd.Run()
	// Output:
	// Version:        1.2.2
	// Go version:     go1.23.4
	// Git commit:     677529b4e7c38e3295e0883a2ee8a536e3a860f4
	// Git tree state: clean
}
