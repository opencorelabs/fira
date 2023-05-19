package auth

import (
	"time"
)

type AccountNamespace string

const (
	AccountNamespaceNone      AccountNamespace = ""
	AccountNamespaceConsumer  AccountNamespace = "consumer"
	AccountNamespaceDeveloper AccountNamespace = "developer"
)

type Account struct {
	ID              string `db:"account_id"`
	Namespace       AccountNamespace
	Valid           bool
	CredentialsType CredentialsType `db:"credentials_type"`
	Credentials     map[string]string

	Name      string
	AvatarURL string `db:"avatar_url"`
	Email     string

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (a *Account) MergeCredentials(c map[string]string) *Account {
	for k, v := range c {
		a.Credentials[k] = v
	}
	return a
}
