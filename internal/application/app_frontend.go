package application

import (
	"bufio"
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
)

func (a *App) StartFrontend(ctx context.Context) error {
	log := a.Logger().Sugar().Named("startup")

	cliDir, cliDirErr := filepath.Abs(a.cfg.ClientDir)
	if cliDirErr != nil {
		return fmt.Errorf("unable to resolve client dir: %w", cliDirErr)
	}

	var yarnArgs []string
	if a.cfg.Debug {
		yarnArgs = append(yarnArgs, "workspace", "@fira/app", "dev")
	} else {
		yarnArgs = append(yarnArgs, "run", "start")
	}

	feUrl, feUrlErr := url.Parse(a.cfg.FrontendUrl)
	if feUrlErr != nil {
		return fmt.Errorf("unable to parse frontend url '%s': %w", a.cfg.FrontendUrl, feUrlErr)
	}

	// start the next subprocess
	a.StartService(ctx, "next-server", func(ctx context.Context, errChan chan error) Finalizer {
		cmd := exec.CommandContext(ctx, "yarn", yarnArgs...)
		cmd.Dir = cliDir
		cmd.Env = os.Environ()

		a.logger.Named("next-process").Info("next server env", zap.Strings("env", cmd.Env))

		go func() {
			defer a.PanicRecovery(errChan)

			logger := a.Logger().Named("next-process")

			stdout, stdoutPipeErr := cmd.StdoutPipe()
			if stdoutPipeErr != nil {
				errChan <- fmt.Errorf("error getting stdout pipe: %w", stdoutPipeErr)
				return
			}

			stderr, stderrPipeErr := cmd.StderrPipe()
			if stderrPipeErr != nil {
				errChan <- fmt.Errorf("error getting stderr pipe: %w", stderrPipeErr)
				return
			}

			go a.pipeToLogger(stdout, zap.InfoLevel, logger, errChan)
			go a.pipeToLogger(stderr, zap.WarnLevel, logger, errChan)

			log.Infow("starting next server")

			errChan <- cmd.Run()
		}()

		return func(ctx context.Context) error {
			if cmd.Process != nil {
				return cmd.Process.Kill()
			}
			return nil
		}
	})

	// proxy requests to the next subprocess
	proxy := httputil.NewSingleHostReverseProxy(feUrl)

	a.mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		proxy.ServeHTTP(writer, request)
	})

	return nil
}

func (a *App) pipeToLogger(closer io.ReadCloser, level zapcore.Level, logger *zap.Logger, errChan chan error) {
	defer a.PanicRecovery(errChan)

	scanner := bufio.NewScanner(closer)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		msg := scanner.Text()
		logger.Log(level, msg)
	}
}
