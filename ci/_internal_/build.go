package ci

import (
	"context"
	"fmt"

	"dagger.io/dagger"
)

func Build(ctx context.Context, client *dagger.Client) error {
	fmt.Println("Building with Dagger")

	// expose host service on port 3306
	hostSrv := client.Host().Service([]dagger.PortForward{
		{Frontend: 7005, Backend: 7005},
	})

	// get reference to the local project
	project := client.Host().Directory(".")

	// define the application build command
	serverEntrypoint := "./src/main.go"

	// get `golang` image
	builder := client.Container().
		From("golang:1.21.7-alpine3.19").
		WithDirectory("/src", project).
		WithWorkdir("/src").
		WithServiceBinding("builder", hostSrv).
		WithEnvVariable("CGO_ENABLED", "0").
		WithExec([]string{"go", "build", serverEntrypoint})

	runner := client.Container().
		From("alpine").
		WithDirectory(".", project).
		WithFile("/bin/main", builder.File("/src/main")).
		WithExec([]string{"chmod", "+x", "/bin/main"}).
		WithEntrypoint([]string{"/bin/main"}).
		WithExposedPort(7005)

	// use secret for registry authentication
	secret := client.SetSecret("password", "!19Riju97!")
	addr, err := runner.WithRegistryAuth("docker.io", "foxbat007", secret).
		Publish(ctx, "foxbat007/taskmanager")
	if err != nil {
		panic(err)
	}
	fmt.Println("Published at:", addr)
	return nil
}
