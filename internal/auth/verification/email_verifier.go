package verification

import (
	"context"
	"fmt"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/email"
	"math/rand"
	"time"
)

type EmailVerifier struct {
	emailer  email.Provider
	accounts auth.AccountStoreProvider
}

func NewEmailVerifier(storeProvider auth.AccountStoreProvider, emailProvider email.Provider) Verifier {
	return &EmailVerifier{
		accounts: storeProvider,
		emailer:  emailProvider,
	}
}

func (l *EmailVerifier) SendVerificationToken(ctx context.Context, account *auth.Account) (map[string]string, error) {
	tok := fmt.Sprintf("%d", rand.Int63n(9999999))
	err := l.emailer.Sender().SendOne(ctx, "auth@mg.opencorelabs.org", account.Credentials["email"], "Verify your email", "email-verification", map[string]string{
		"verification_token": tok,
	})
	if err != nil {
		return nil, fmt.Errorf("error sending email: %w", err)
	}
	return map[string]string{
		"email_verification_token":           tok,
		"email_verification_token_timestamp": time.Now().Format(time.RFC3339Nano),
	}, nil
}

func (l *EmailVerifier) VerifyWithToken(ctx context.Context, namespace auth.AccountNamespace, token string) (*auth.Account, error) {
	cred := map[string]string{
		"email_verification_token": token,
	}
	user, userErr := l.accounts.AccountStore().FindByCredentials(ctx, namespace, cred)
	if userErr != nil {
		return nil, auth.ErrNoAccount
	}

	user.Valid = true
	user.Credentials["email_verification_token"] = ""
	user.Credentials["email_verification_token_timestamp"] = ""

	err := l.accounts.AccountStore().Update(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error updating account: %w", err)
	}
	return user, nil
}
