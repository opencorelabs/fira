package email_password

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/auth/verification"
	"golang.org/x/crypto/bcrypt"
)

type Data struct {
	Email    string
	Password string
}

type EmailPasswordBackend struct {
	storeProvider        auth.AccountStoreProvider
	verificationProvider verification.Provider
}

func New(storeProvider auth.AccountStoreProvider, verifProvider verification.Provider) auth.Backend {
	return &EmailPasswordBackend{
		storeProvider:        storeProvider,
		verificationProvider: verifProvider,
	}
}

func (e *EmailPasswordBackend) Register(ctx context.Context, namespace auth.AccountNamespace, credentials map[string]string) (*auth.Account, error) {
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
		Valid:           false, // account is initially invalid, email must be verified
		Namespace:       namespace,
		CredentialsType: auth.CredentialsTypeEmailPassword,
		Credentials: map[string]string{
			"email":    d.Email,
			"password": fmt.Sprintf("%x", pw),
		},
		Email: d.Email,
	}

	if saveErr := e.storeProvider.AccountStore().Create(ctx, acct, uniqueCredentials); saveErr != nil {
		return nil, fmt.Errorf("failed to save account: %w", saveErr)
	}

	// send verification token via email
	additionalCredentials, sendErr := e.verificationProvider.Email().SendVerificationToken(ctx, acct)
	if sendErr != nil {
		return nil, fmt.Errorf("failed to send verification token: %w", sendErr)
	}

	// save verification token to account
	updatedAcct := acct.MergeCredentials(additionalCredentials)
	if saveErr := e.storeProvider.AccountStore().Update(ctx, updatedAcct); saveErr != nil {
		return nil, fmt.Errorf("failed to save verification token: %w", saveErr)
	}

	return acct, nil
}

func (e *EmailPasswordBackend) Authenticate(ctx context.Context, namespace auth.AccountNamespace, credentials map[string]string) (*auth.Account, error) {
	d, dErr := decode(credentials)
	if dErr != nil {
		return nil, dErr
	}

	acct, acctErr := e.storeProvider.AccountStore().FindByCredentials(ctx, namespace, map[string]string{
		"email": d.Email,
	})

	if acctErr != nil {
		return nil, auth.ErrNoAccount
	}

	if acct.CredentialsType != auth.CredentialsTypeEmailPassword {
		return nil, auth.ErrInvalidCredentials
	}

	decodedPw, decodeErr := hex.DecodeString(acct.Credentials["password"])
	if decodeErr != nil {
		return nil, fmt.Errorf("failed to decode password: %w", decodeErr)
	}

	if err := bcrypt.CompareHashAndPassword(decodedPw, []byte(d.Password)); err != nil {
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
