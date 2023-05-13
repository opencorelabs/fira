package verification

import (
	"context"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/logging"
)

type Verifier interface {
	SendVerificationToken(ctx context.Context, account *auth.Account) (map[string]string, error)
	VerifyWithToken(ctx context.Context, token string) (*auth.Account, error)
}

type Provider interface {
	Email() Verifier
}

type DefaultProvider struct {
	loggingProvider      logging.Provider
	accountStoreProvider auth.AccountStoreProvider
}

func NewDefaultProvider(loggingProvider logging.Provider, accountStoreProvider auth.AccountStoreProvider) Provider {
	return &DefaultProvider{
		loggingProvider:      loggingProvider,
		accountStoreProvider: accountStoreProvider,
	}
}

func (d *DefaultProvider) Email() Verifier {
	return NewLoggingVerifier(d.loggingProvider, d.accountStoreProvider)
}
