package main

import (
	"context"
	"fmt"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/opencorelabs/fira/internal/application"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"os"
)

func main() {
	app, appErr := application.NewApp()
	if appErr != nil {
		fmt.Println("app init error:", appErr)
		os.Exit(1)
	}

	cfg := app.Config()

	cliApp := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "serve",
				Usage: "Run the server",
				Action: func(c *cli.Context) error {
					ctx, cancel := context.WithCancel(context.Background())
					log := app.Logger().Named("startup").Sugar()

					app.StartDB(ctx)

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

					return nil
				},
			},

			{
				Name:  "bootstrap",
				Usage: "Bootstrap the embedded database environment",
				Action: func(c *cli.Context) error {
					postgres := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().
						Username(cfg.EmbeddedPostgres.Username).
						Password(cfg.EmbeddedPostgres.Password).
						Database(cfg.EmbeddedPostgres.Database).
						Port(uint32(cfg.EmbeddedPostgres.Port)).
						Version(embeddedpostgres.V15).
						RuntimePath(cfg.EmbeddedPostgres.RuntimePath).
						DataPath(cfg.EmbeddedPostgres.DataPath).
						BinariesPath(cfg.EmbeddedPostgres.BinariesPath),
					)
					err := postgres.Start()

					if err != nil {
						return fmt.Errorf("postgres init failed: %w", err)
					}

					if stopErr := postgres.Stop(); stopErr != nil {
						return fmt.Errorf("postgres stop failed: %w", stopErr)
					}

					return nil
				},
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		app.Logger().Fatal("cli app run failed", zap.Error(err))
	}
}
