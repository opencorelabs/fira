package email_password

import (
	"context"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/opencorelabs/fira/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

type Data struct {
	Email    string
	Password string
}

type EmailPasswordBackend struct {
	storeProvider auth.AccountStoreProvider
}

func New(storeProvider auth.AccountStoreProvider) auth.Backend {
	return &EmailPasswordBackend{
		storeProvider: storeProvider,
	}
}

func (e *EmailPasswordBackend) Register(ctx context.Context, credentials map[string]string) (*auth.Account, error) {
	d, dErr := decode(credentials)
	if dErr != nil {
		return nil, dErr
	}

	pw, pwErr := bcrypt.GenerateFromPassword([]byte(d.Password), bcrypt.DefaultCost)
	if pwErr != nil {
		return nil, fmt.Errorf("failed to hash password: %w", pwErr)
	}

	uniqueCredentials := map[string]string{"email": d.Email}
	acct := &auth.Account{
		// TODO: account should initially be invalid
		Valid:           true, // account is initially invalid, email must be verified
		CredentialsType: auth.CredentialsTypeEmailPassword,
		Credentials: map[string]string{
			"email":    d.Email,
			"password": fmt.Sprintf("%x", pw),
		},
	}

	if saveErr := e.storeProvider.AccountStore().Create(ctx, acct, uniqueCredentials); saveErr != nil {
		return nil, fmt.Errorf("failed to save account: %w", saveErr)
	}

	return acct, nil
}

func (e *EmailPasswordBackend) Authenticate(ctx context.Context, credentials map[string]string) (*auth.Account, error) {
	d, dErr := decode(credentials)
	if dErr != nil {
		return nil, dErr
	}

	acct, acctErr := e.storeProvider.AccountStore().FindByCredentials(ctx, map[string]string{
		"email": d.Email,
	})

	if acctErr != nil {
		return nil, auth.ErrNoAccount
	}

	if acct.CredentialsType != auth.CredentialsTypeEmailPassword {
		return nil, auth.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(acct.Credentials["password"]), []byte(d.Password)); err != nil {
		return nil, auth.ErrInvalidCredentials
	}

	return acct, nil
}

func decode(m map[string]string) (*Data, error) {
	d := &Data{}
	if err := mapstructure.Decode(m, d); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}
	return d, nil
}
