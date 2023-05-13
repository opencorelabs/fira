package logging

import (
	"fmt"
	"github.com/opencorelabs/fira/internal/config"
	"go.uber.org/zap"
)

type Provider interface {
	Logger() *zap.Logger
}

func Init(cfg config.Provider) (*zap.Logger, error) {
	var logger *zap.Logger
	var loggerErr error
	if cfg.Config().Debug {
		logger, loggerErr = zap.NewDevelopment()
	} else {
		logger, loggerErr = zap.NewProduction()
	}
	if loggerErr != nil {
		return nil, fmt.Errorf("unable to init logger: %w", loggerErr)
	}
	logger = logger.WithOptions(zap.AddStacktrace(zap.FatalLevel))
	return logger, nil
}
