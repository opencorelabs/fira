package developer

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AppStore interface {
	CreateApp(ctx context.Context, app *App) error
	UpdateApp(ctx context.Context, app *App) error
	GetAppsByAccountID(ctx context.Context, accountID string) ([]*App, error)
	GetAppByID(ctx context.Context, appID string) (*App, error)
}

type AppStoreProvider interface {
	AppStore() AppStore
}

type AppClaims struct {
	jwt.RegisteredClaims
	AppID       string `json:"app_id"`
	Environment string `json:"env"`
}

type App struct {
	ID        string `db:"app_id"`
	Name      string
	AccountID string `db:"account_id"`
	Tokens    TokenMap
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (a *App) JWT(env Environment) (string, error) {
	tokens := a.Tokens[env]
	if len(tokens) == 0 {
		return "", fmt.Errorf("no tokens for environment: %s", env)
	}
	tok := tokens[len(tokens)-1]
	jwtTok := jwt.NewWithClaims(jwt.SigningMethodHS512, AppClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "fira",
			ExpiresAt: jwt.NewNumericDate(tok.ValidUntil),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		AppID:       a.ID,
		Environment: string(env),
	})
	str, err := jwtTok.SignedString(tok.Key)
	if err != nil {
		return "", fmt.Errorf("failed to sign jwt: %w", err)
	}
	return str, nil
}

func (a *App) Rotate(env Environment) error {
	if a.Tokens == nil {
		a.Tokens = make(TokenMap)
	}
	tok, tokGenErr := generateToken(env)
	if tokGenErr != nil {
		return fmt.Errorf("failed to generate token: %w", tokGenErr)
	}
	a.Tokens[env] = append(a.Tokens[env], tok)
	return nil
}

func (a *App) Invalidate(jwt string) error {
	return fmt.Errorf("not implemented")
}

func generateToken(env Environment) (tok Token, err error) {
	expiry, hasExpiry := TokenExpiryMap[env]
	if !hasExpiry {
		err = fmt.Errorf("no expiry for environment: %s", env)
		return
	}
	randBytes := make([]byte, 32)
	if _, err = rand.Read(randBytes); err != nil {
		err = fmt.Errorf("failed to generate random bytes: %w", err)
		return
	}
	tok = Token{
		Key:        randBytes,
		ValidUntil: time.Now().Add(expiry),
	}
	return
}
