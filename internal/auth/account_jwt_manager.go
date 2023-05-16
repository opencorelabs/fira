package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/opencorelabs/fira/internal/logging"
	"go.uber.org/zap"
	"time"
)

type AccountJWTManager struct {
	secrets              [][]byte
	duration             time.Duration
	accountStoreProvider AccountStoreProvider
	logger               *zap.SugaredLogger
}

func NewAccountJWTManager(
	secrets [][]byte,
	duration time.Duration,
	loggingProvider logging.Provider,
	accountStoreProvider AccountStoreProvider,
) JWTManager {
	return &AccountJWTManager{
		secrets:              secrets,
		duration:             duration,
		accountStoreProvider: accountStoreProvider,
		logger:               loggingProvider.Logger().Named("account-jwt-provider").Sugar(),
	}
}

func (a *AccountJWTManager) Generate(ctx context.Context, principal interface{}) (string, error) {
	account, ok := principal.(*Account)
	if !ok {
		return "", fmt.Errorf("principal is not an account")
	}
	claims := FiraClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "fira",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		AccountID:        account.ID,
		AccountNamespace: string(account.Namespace),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(a.secrets[len(a.secrets)-1])
}

func (a *AccountJWTManager) Verify(ctx context.Context, tokenStr string) (context.Context, error) {
	claims := &FiraClaims{}
	var token *jwt.Token
	var tokenErr error
	var valid bool

	// iterate secrets in reverse order to find the correct one
	// as new secrets are added, they are added to the end of the slice
	for i := len(a.secrets) - 1; i >= 0; i-- {
		token, tokenErr = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return a.secrets[i], nil
		})
		if tokenErr != nil {
			if errors.Is(tokenErr, jwt.ErrTokenSignatureInvalid) {
				a.logger.Debugw("invalid signature", "secretID", i)
			}
			return nil, fmt.Errorf("failed to parse token: %w", tokenErr)
		}
		if token.Valid {
			valid = true
			break
		}
	}

	if !valid {
		return nil, ErrInvalidToken
	}

	var namespace AccountNamespace
	if claims.AccountNamespace == string(AccountNamespaceConsumer) {
		namespace = AccountNamespaceConsumer
	} else if claims.AccountNamespace == string(AccountNamespaceDeveloper) {
		namespace = AccountNamespaceDeveloper
	} else {
		a.logger.Debugw("invalid namespace", "namespace", claims.AccountNamespace)
		return nil, StandardRejectionCode
	}

	account, accountErr := a.accountStoreProvider.AccountStore().FindAccountByID(ctx, namespace, claims.AccountID)
	if accountErr != nil {
		a.logger.Debugw("account not found", "account_id", claims.AccountID)
		return nil, StandardRejectionCode
	}

	if !account.Valid {
		a.logger.Debugw("account is not valid", "account_id", claims.AccountID)
		return nil, StandardRejectionCode
	}

	return context.WithValue(ctx, firaAccountKey, account), nil
}
