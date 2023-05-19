package verification

import (
	"context"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/email"
	"github.com/opencorelabs/fira/internal/logging"
)

type Verifier interface {
	SendVerificationToken(ctx context.Context, account *auth.Account) (map[string]string, error)
	VerifyWithToken(ctx context.Context, namespace auth.AccountNamespace, token string) (*auth.Account, error)
}

type Provider interface {
	Email() Verifier
}

type DefaultProvider struct {
	loggingProvider      logging.Provider
	accountStoreProvider auth.AccountStoreProvider
	emailProvider        email.Provider
}

func NewDefaultProvider(loggingProvider logging.Provider, accountStoreProvider auth.AccountStoreProvider, emailProvider email.Provider) Provider {
	return &DefaultProvider{
		loggingProvider:      loggingProvider,
		accountStoreProvider: accountStoreProvider,
		emailProvider:        emailProvider,
	}
}

func (d *DefaultProvider) Email() Verifier {
	return NewEmailVerifier(d.accountStoreProvider, d.emailProvider)
}
