package main

import (
	"fmt"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/opencorelabs/fira/internal/application"
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

	postgres := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().
		Username(cfg.EmbeddedPostgres.Username).
		Password(cfg.EmbeddedPostgres.Password).
		Database(cfg.EmbeddedPostgres.Database).
		Port(uint32(cfg.EmbeddedPostgres.Port)).
		Version(embeddedpostgres.V15).
		RuntimePath("/tmp/pg-runtime").
		BinariesPath(cfg.EmbeddedPostgres.BinariesPath),
	)
	err := postgres.Start()

	if err != nil {
		app.Logger().Fatal("postgres init failed", zap.Error(err))
	}

	if stopErr := postgres.Stop(); stopErr != nil {
		app.Logger().Fatal("postgres stop failed", zap.Error(stopErr))
	}
}
