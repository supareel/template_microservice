package ci

import (
	"context"
	"fmt"

	"dagger.io/dagger"
)

func Test(ctx context.Context, client *dagger.Client) error {
	fmt.Println("Testing with Dagger")

	// get reference to the local project
	project := client.Host().Directory(".")

	// Database service used for application tests
	database := client.Container().From("postgres:15.2").
		WithEnvVariable("POSTGRES_PASSWORD", "test").
		WithExec([]string{"postgres"}).
		WithExposedPort(5432).
		AsService()

	// Run application tests
	out, err := client.Container().From("golang:1.21.7-alpine3.19").
		WithServiceBinding("db", database).     // bind database with the name db
		WithEnvVariable("DB_HOST", "db").       // db refers to the service binding
		WithEnvVariable("DB_PASSWORD", "test"). // password set in db container
		WithEnvVariable("DB_USER", "postgres"). // default user in postgres image
		WithEnvVariable("DB_NAME", "postgres"). // default db name in postgres image
		WithDirectory("/src", project).
		WithWorkdir("/src").
		WithEnvVariable("CGO_ENABLED", "0").
		WithExec([]string{"go", "test", "-v", "-cover", "./..."}).
		Stdout(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Print(out)
	fmt.Println("Tests passed successfully")

	return nil
}
