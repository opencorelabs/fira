package in_memory

import (
	"context"
	"github.com/google/uuid"
	"github.com/opencorelabs/fira/internal/auth"
	cmap "github.com/orcaman/concurrent-map/v2"
)

type InMemoryAccountStore struct {
	data cmap.ConcurrentMap[string, *auth.Account]
}

func New() auth.AccountStore {
	return &InMemoryAccountStore{
		cmap.New[*auth.Account](),
	}
}

func (i *InMemoryAccountStore) FindAccountByID(ctx context.Context, id string) (*auth.Account, error) {
	acct, has := i.data.Get(id)
	if !has {
		return nil, auth.ErrNoAccount
	}
	return acct, nil
}

func (i *InMemoryAccountStore) FindByCredentials(ctx context.Context, creds map[string]string) (*auth.Account, error) {
	for _, acct := range i.data.Items() {
		nMatches := 0
		// match all provided credentials
		for k, v := range creds {
			if v == acct.Credentials[k] {
				nMatches++
			}
		}
		if nMatches == len(creds) {
			return acct, nil
		}
	}
	return nil, auth.ErrNoAccount
}

func (i *InMemoryAccountStore) Create(ctx context.Context, account *auth.Account, creds map[string]string) error {
	existing, _ := i.FindByCredentials(ctx, creds)
	if existing != nil {
		return auth.ErrAccountExists
	}

	acctId := uuid.NewString()
	account.ID = acctId
	i.data.Set(acctId, account)
	return nil
}

func (i *InMemoryAccountStore) Update(ctx context.Context, account *auth.Account) error {
	i.data.Set(account.ID, account)
	return nil
}
