package application

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os/exec"
)

type App struct {
	cfg         *Config
	logger      *zap.Logger
	frontendCmd *exec.Cmd
	mux         *http.ServeMux
}

func NewApp() (*App, error) {
	cfg, cfgErr := InitConfig()
	if cfgErr != nil {
		return nil, fmt.Errorf("unable to init config: %w", cfgErr)
	}

	var logger *zap.Logger
	var loggerErr error
	if cfg.Debug {
		logger, loggerErr = zap.NewDevelopment()
	} else {
		logger, loggerErr = zap.NewProduction()
	}
	if loggerErr != nil {
		return nil, fmt.Errorf("unable to init logger: %w", loggerErr)
	}

	return &App{
		cfg:    cfg,
		logger: logger,
		mux:    http.NewServeMux(),
	}, nil
}

func (a *App) Config() *Config {
	return a.cfg
}

func (a *App) Logger() *zap.Logger {
	return a.logger
}

func (a *App) Mux() *http.ServeMux {
	return a.mux
}
