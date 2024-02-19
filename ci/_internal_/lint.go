package ci

import (
	"context"

	"dagger.io/dagger"
)

func Lint(ctx context.Context, client *dagger.Client) error {

	// Project to test
	src := client.Host().Directory(".")

	// golangci service used for application lint
	golangci := client.Container(dagger.ContainerOpts{Platform: "linux/amd64"}).
		From("docker/golangci-lint:1.55.1-go1.21.6").
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"golangci-lint", "run", "-v"}).
		AsService()

	golangci.Start(ctx)
	defer golangci.Stop(ctx, dagger.ServiceStopOpts{
		Kill: true,
	})
	return nil
}
