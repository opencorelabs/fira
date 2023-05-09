package main

import (
	"context"
	"fmt"
	"github.com/opencorelabs/fira/internal/application"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	app, appErr := application.NewApp()
	if appErr != nil {
		fmt.Println("app init error:", appErr)
		os.Exit(1)
	}
	log := app.Logger().Named("startup").Sugar()

	startFrontendError := app.StartFrontend(ctx)
	if startFrontendError != nil {
		log.Fatalw("unable to start frontend server", "error", startFrontendError)
	}

	startGrpcErr := app.StartGRPC(ctx)
	if startGrpcErr != nil {
		log.Fatalw("unable to start api server", "error", startGrpcErr)
	}

	app.StartHTTP(ctx)

	app.Wait(cancel)
}
