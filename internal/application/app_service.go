package application

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ServiceFn func(ctx context.Context, errChan chan error) Finalizer
type Finalizer func(ctx context.Context) error

// StartService will run a blocking function in a loop until ctx.Done() is signaled
func (a *App) StartService(ctx context.Context, name string, fn ServiceFn) {
	log := a.Logger().Named(name + "-svc").Sugar()
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
				timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				if finalizer != nil {
					if err := finalizer(timeoutCtx); err != nil {
						log.Warnw("error stopping service", "error", err)
					}
				}
				log.Infow("service is terminated")
				return
			}
		}
	}()
}

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
