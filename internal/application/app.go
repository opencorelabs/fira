package application

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type ServiceFn func(ctx context.Context, errChan chan error) Finalizer
type Finalizer func(ctx context.Context) error

// App is a service container and supervision tree for Fira. StartService() can be used to fork a new
// "process" (blocking function) which will be supervised and restarted if it fails.
type App struct {
	cfg    *Config
	logger *zap.Logger
	mux    *http.ServeMux
	wg     *sync.WaitGroup
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

	logger.Sugar().Named("startup").Infow("config initialized", "debug", cfg.Debug)

	return &App{
		cfg:    cfg,
		logger: logger,
		mux:    http.NewServeMux(),
		wg:     &sync.WaitGroup{},
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

// StartService will run a blocking function in a loop until ctx.Done() is signaled
func (a *App) StartService(ctx context.Context, name string, fn ServiceFn) {
	log := a.Logger().Named(name + "-supervisor").Sugar()
	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		errChan := make(chan error)
		for {
			finalizer := fn(ctx, errChan)
			select {
			case err := <-errChan:
				log.Errorw("service failed, restarting", "error", err)
				time.Sleep(300 * time.Millisecond)
			case <-ctx.Done():
				log.Infow("finalizing service")
				// encapsulate in an anonymous function to properly cancel the timout context
				func() {
					timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()
					if finalizer != nil {
						if err := finalizer(timeoutCtx); err != nil {
							log.Warnw("error stopping service", "error", err)
						}
					}
				}()
				log.Infow("service is terminated")
				return
			}
		}
	}()
}

func (a *App) PanicRecovery(errChan chan error) {
	if p := recover(); p != nil {
		errChan <- fmt.Errorf("recovered from panic: %#v", p)
	}
}

// Wait will block until a kill signal is received to the process, after which it will call the provided
// context.CancelFunc and wait for all services started via StartService to exit.
func (a *App) Wait(cancelFunc context.CancelFunc) {
	log := a.logger.Named("startup").Sugar()
	log.Info("waiting for shutdown signal [sigint, sigterm]")
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Info("shutdown signal received")
	cancelFunc()
	a.wg.Wait()
}