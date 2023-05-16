package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
)

var firaAccountKey = struct{}{}

type FiraClaims struct {
	jwt.RegisteredClaims
	// for account tokens
	AccountID        string `json:"actid,omitempty"`
	AccountNamespace string `json:"actns,omitempty"`
	// for app tokens
	AppID          string `json:"appid,omitempty"`
	AppEnvironment string `json:"appenv,omitempty"`
}

type JWTManager interface {
	// Generate returns a JWT string
	Generate(ctx context.Context, principal interface{}) (string, error)

	// Verify returns a context with the principal set
	Verify(ctx context.Context, tokenStr string) (context.Context, error)
}
