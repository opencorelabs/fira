package developer

import (
	"crypto/rand"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Environment string

const (
	EnvironmentDevelopment Environment = "development"
	EnvironmentStaging     Environment = "staging"
	EnvironmentProduction  Environment = "production"
)

var (
	TokenExpiryMap = map[Environment]time.Duration{
		EnvironmentDevelopment: time.Hour * 24 * 30,
		EnvironmentStaging:     time.Hour * 24 * 90,
		EnvironmentProduction:  time.Hour * 24 * 365,
	}
)

type AppClaims struct {
	jwt.RegisteredClaims
	AppID       string `json:"app_id"`
	Environment string `json:"env"`
}

type Token struct {
	Key        []byte
	ValidUntil time.Time
}

type App struct {
	ID        string
	Name      string
	AccountID string
	Tokens    map[Environment][]Token
}

func (a *App) JWT(env Environment) (string, error) {
	tokens := a.Tokens[env]
	if len(tokens) == 0 {
		newTok, tokErr := generateToken(env)
		if tokErr != nil {
			return "", fmt.Errorf("failed to generate token: %w", tokErr)
		}
		a.Tokens[env] = append(a.Tokens[env], newTok)
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
	tok, tokGenErr := generateToken(env)
	if tokGenErr != nil {
		return fmt.Errorf("failed to generate token: %w", tokGenErr)
	}
	a.Tokens[env] = append(a.Tokens[env], tok)
	return nil
}

type AppStore interface {
	CreateApp(app *App) error
	GetAppsByAccountID(accountID string) ([]*App, error)
	GetAppByJWT(token string) (*App, error)
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
