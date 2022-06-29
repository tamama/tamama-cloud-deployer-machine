package main

import (
	"context"
	"log"
	"os"

	"github.com/tamama/tamama-cloud-deployer-machine/pkg/application"
)

func main() {
	ctx := context.Background()
	error_from_application := make(chan error)

	app := application.GetApplication(ctx, "v1")
	go app.Run(ctx, error_from_application)

	for {
		err := <-error_from_application
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}