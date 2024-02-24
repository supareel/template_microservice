package ci

import (
	"context"
	"fmt"

	"dagger.io/dagger"
)

func Test(ctx context.Context, client *dagger.Client) error {
	fmt.Println("Building with Dagger")

	// expose host service on port 3306
	hostSrv := client.Host().Service([]dagger.PortForward{
		{Frontend: 7001, Backend: 7001},
	})

	// get reference to the local project
	project := client.Host().Directory(".")

	// define the application build command
	path := "build/"
	serverEntrypoint := "src/cmd/server/server.go"

	// get `golang` image
	builder := client.Container().
		From("golang:latest").
		WithDirectory("/src", project).
		WithWorkdir("/src").
		WithServiceBinding("db", hostSrv).
		WithEnvVariable("CGO_ENABLED", "0").
		WithExec([]string{"go", "test", "-v", " -cover", "./src/..."})

	fmt.Println("Tests passed successfully")

	prodImage := client.Container().
		From("alpine").
		WithFile(path, builder.File(serverEntrypoint)).
		WithEntrypoint([]string{path}).WithExposedPort(7001)

	addr, err := prodImage.Publish(ctx, "hub.docker.io/multistage")
	if err != nil {
		panic(err)
	}

	fmt.Println(addr)

	return nil
}
