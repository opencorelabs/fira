package developer

import (
	"context"
	"time"
)

type Environment string

const (
	EnvironmentSandbox     Environment = "sandbox"
	EnvironmentDevelopment Environment = "development"
	EnvironmentProduction  Environment = "production"
)

var (
	TokenExpiryMap = map[Environment]time.Duration{
		EnvironmentSandbox:     time.Hour * 24 * 90,
		EnvironmentDevelopment: time.Hour * 24 * 30,
		EnvironmentProduction:  time.Hour * 24 * 365,
	}
	appKey     = struct{}{}
	environKey = struct{}{}
)

func WithEnvironment(ctx context.Context, env Environment) context.Context {
	return context.WithValue(ctx, environKey, env)
}

func EnvironmentFromContext(ctx context.Context) (env Environment, has bool) {
	val := ctx.Value(environKey)
	env, has = val.(Environment)
	return
}

func AppFromContext(ctx context.Context) (app *App, has bool) {
	val := ctx.Value(appKey)
	app, has = val.(*App)
	return
}
