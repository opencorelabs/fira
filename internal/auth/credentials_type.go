package auth

import (
	"database/sql/driver"
	"fmt"
)

type CredentialsType int

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
