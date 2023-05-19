package auth

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type CredentialsType int
type AccountNamespace string

const (
	CredentialsTypeNone CredentialsType = iota
	CredentialsTypeEmailPassword
	CredentialsTypeOAuth
)

func (c *CredentialsType) Scan(src interface{}) error {
	str, ok := src.(string)
	if !ok {
		return fmt.Errorf("scan source was not []bytes got %T", src)
	}
	switch str {
	case "none":
		*c = CredentialsTypeNone
	case "email_password":
		*c = CredentialsTypeEmailPassword
	case "oauth":
		*c = CredentialsTypeOAuth
	default:
		return fmt.Errorf("invalid CredentialsType %s", str)
	}
	return nil
}

func (c CredentialsType) Value() (driver.Value, error) {
	switch c {
	case CredentialsTypeNone:
		return "none", nil
	case CredentialsTypeEmailPassword:
		return "email_password", nil
	case CredentialsTypeOAuth:
		return "oauth", nil
	default:
		return nil, fmt.Errorf("invalid CredentialsType %v", c)
	}
}

const (
	AccountNamespaceNone      AccountNamespace = ""
	AccountNamespaceConsumer  AccountNamespace = "consumer"
	AccountNamespaceDeveloper AccountNamespace = "developer"
)

type Account struct {
	ID              string
	Namespace       AccountNamespace
	Valid           bool
	CredentialsType CredentialsType
	Credentials     map[string]string

	Name      string
	AvatarURL string
	Email     string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *Account) MergeCredentials(c map[string]string) *Account {
	for k, v := range c {
		a.Credentials[k] = v
	}
	return a
}
