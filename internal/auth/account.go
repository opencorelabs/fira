package auth

import "time"

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
