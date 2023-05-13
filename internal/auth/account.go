package auth

type Account struct {
	ID              string
	Valid           bool
	CredentialsType CredentialsType
	Credentials     map[string]string
}

func (a *Account) MergeCredentials(c map[string]string) *Account {
	for k, v := range c {
		a.Credentials[k] = v
	}
	return a
}
