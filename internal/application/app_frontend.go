package application

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
)

func (a *App) StartFrontend(ctx context.Context) error {
	cliDir, cliDirErr := filepath.Abs(a.cfg.ClientDir)
	if cliDirErr != nil {
		return fmt.Errorf("unable to resolve client dir: %w", cliDirErr)
	}

	yarnArgs := []string{"run"}
	if a.cfg.Debug {
		yarnArgs = append(yarnArgs, "dev")
	} else {
		yarnArgs = append(yarnArgs, "start")
	}

	cmd := exec.CommandContext(ctx, "yarn", yarnArgs...)
	cmd.Dir = cliDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("yarn failed to start: %w", err)
	}

	a.frontendCmd = cmd

	return a.ServeFrontend()
}

func (a *App) StopFrontend() error {
	if a.frontendCmd != nil && a.frontendCmd.Process != nil {
		if killErr := a.frontendCmd.Process.Kill(); killErr != nil {
			return fmt.Errorf("failed to kill frontend process: %w", killErr)
		}
	}
	return nil
}

func (a *App) ServeFrontend() error {
	feUrl, feUrlErr := url.Parse(a.cfg.FrontendURL)
	if feUrlErr != nil {
		return fmt.Errorf("unable to parse frontend url '%s': %w", a.cfg.FrontendURL, feUrlErr)
	}

	proxy := httputil.NewSingleHostReverseProxy(feUrl)

	a.mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		proxy.ServeHTTP(writer, request)
	})

	return nil
}
