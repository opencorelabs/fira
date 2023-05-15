package auth

import (
	"context"
	"errors"
)

type CredentialsType int
type AccountNamespace string

const (
	CredentialsTypeNone CredentialsType = iota
	CredentialsTypeEmailPassword
	CredentialsTypeOAuth
)

const (
	AccountNamespaceNone      AccountNamespace = ""
	AccountNamespaceConsumer  AccountNamespace = "consumer"
	AccountNamespaceDeveloper AccountNamespace = "developer"
)

var (
	ErrNoAccount          = errors.New("no account found")
	ErrAccountExists      = errors.New("account already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnknownAuthorizer  = errors.New("unknown authorizer")
)

type AccountStore interface {
	FindAccountByID(ctx context.Context, namespace AccountNamespace, id string) (*Account, error)
	FindByCredentials(ctx context.Context, namespace AccountNamespace, creds map[string]string) (*Account, error)
	Create(ctx context.Context, account *Account, creds map[string]string) error
	Update(ctx context.Context, account *Account) error
}

type AccountStoreProvider interface {
	AccountStore() AccountStore
}

type Backend interface {
	// Register creates a new account from the provided credentials. If an account
	// already exists, returns ErrAccountExists.
	Register(ctx context.Context, namespace AccountNamespace, credentials map[string]string) (*Account, error)

	// Authenticate attempts to authenticate an account using the provided credentials.
	// The credentials are expected to be in the format of the CredentialsType associated
	// with the Backend.
	Authenticate(ctx context.Context, namespace AccountNamespace, credentials map[string]string) (*Account, error)
}

type Registry interface {
	// RegisterBackend registers a backend for the given CredentialsType.
	RegisterBackend(ct CredentialsType, backend Backend)

	// GetBackend returns the backend for the given CredentialsType. If there is no
	// backend registered for the given CredentialsType, returns ErrUnknownAuthorizer.
	GetBackend(ct CredentialsType) (Backend, error)
}
