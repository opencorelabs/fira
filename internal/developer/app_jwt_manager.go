package developer

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/logging"
	"go.uber.org/zap"
	"time"
)

type AppJWTManager struct {
	storeProvider AppStoreProvider
	logger        *zap.SugaredLogger
}

func NewAppJWTManager(logProvider logging.Provider, appStoreProvider AppStoreProvider) auth.JWTManager {
	return &AppJWTManager{
		storeProvider: appStoreProvider,
		logger:        logProvider.Logger().Named("app-jwt-manager").Sugar(),
	}
}

func (a *AppJWTManager) Generate(ctx context.Context, principal interface{}) (string, error) {
	app, isApp := principal.(*App)
	if !isApp {
		return "", fmt.Errorf("principal is not an app")
	}

	env, hasEng := EnvironmentFromContext(ctx)
	if !hasEng {
		return "", errors.New("env not found in context")
	}

	if len(app.Tokens[env]) == 0 {
		return "", fmt.Errorf("no tokens found for app %s in env %s", app.ID, env)
	}

	tok := app.Tokens[env][len(app.Tokens[env])-1]
	jwtTok := jwt.NewWithClaims(jwt.SigningMethodHS512, AppClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "fira",
			ExpiresAt: jwt.NewNumericDate(tok.ValidUntil),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		AppID:       app.ID,
		Environment: string(env),
	})
	str, err := jwtTok.SignedString(tok.Key)
	if err != nil {
		return "", fmt.Errorf("failed to sign jwt: %w", err)
	}

	return str, nil
}

func (a *AppJWTManager) Verify(ctx context.Context, tokenStr string) (context.Context, error) {
	claims := &AppClaims{}
	_, _, parseErr := jwt.NewParser().ParseUnverified(tokenStr, claims)
	if parseErr != nil {
		return nil, fmt.Errorf("failed to parse token: %w", parseErr)
	}
	if claims.AppID == "" {
		return nil, fmt.Errorf("app_id claim is empty")
	}
	if claims.Environment == "" {
		return nil, fmt.Errorf("env claim is empty")
	}

	app, appErr := a.storeProvider.AppStore().GetAppByID(claims.AppID)
	if appErr != nil {
		return nil, fmt.Errorf("failed to get app by id: %w", appErr)
	}

	environ := Environment(claims.Environment)
	_, hasEnv := TokenExpiryMap[environ]
	if !hasEnv {
		return nil, fmt.Errorf("invalid env claim")
	}

	appTokens := app.Tokens[environ]
	var token *jwt.Token
	var tokenErr error
	var valid bool

	// iterate tokens in reverse to find the right key
	for i := len(appTokens) - 1; i >= 0; i-- {
		token, tokenErr = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return appTokens[i].Key, nil
		})
		if tokenErr != nil {
			if errors.Is(tokenErr, jwt.ErrTokenSignatureInvalid) {
				a.logger.Debugw("invalid signature", "secretID", i)
				continue
			}
			return nil, fmt.Errorf("failed to parse token: %w", tokenErr)
		}
		if token.Valid {
			valid = true
			break
		}
	}

	if !valid {
		return nil, auth.ErrInvalidToken
	}

	environCtx := context.WithValue(ctx, environKey, environ)

	return context.WithValue(environCtx, appKey, app), nil
}
