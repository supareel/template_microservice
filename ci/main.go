package main

import (
	"context"
	"flag"
	ci "gomicro/ci/_internal_"
	"os"

	"dagger.io/dagger"
)

func main() {
	var (
		buildFlag bool
		lintFlag  bool
	)

	flag.BoolVar(&buildFlag, "build", false, "used to run build command")
	flag.BoolVar(&lintFlag, "lint", false, "used to run build command")

	flag.Parse()

	// initialize Dagger client
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}

	defer client.Close()

	if buildFlag {
		if err := ci.Build(ctx, client); err != nil {
			panic(err)
		}
	}

	if lintFlag {
		if err := ci.Lint(ctx, client); err != nil {
			panic(err)
		}
	}

}
