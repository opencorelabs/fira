package in_memory

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/opencorelabs/fira/internal/auth"
	cmap "github.com/orcaman/concurrent-map/v2"
)

type InMemoryAccountStore struct {
	data map[auth.AccountNamespace]cmap.ConcurrentMap[string, *auth.Account]
}

func New() auth.AccountStore {
	return &InMemoryAccountStore{
		map[auth.AccountNamespace]cmap.ConcurrentMap[string, *auth.Account]{
			auth.AccountNamespaceDeveloper: cmap.New[*auth.Account](),
			auth.AccountNamespaceConsumer:  cmap.New[*auth.Account](),
		},
	}
}

func (i *InMemoryAccountStore) FindAccountByID(ctx context.Context, namespace auth.AccountNamespace, id string) (*auth.Account, error) {
	if _, hasNs := i.data[namespace]; !hasNs {
		return nil, fmt.Errorf("invalid namespace: %s", namespace)
	}
	acct, has := i.data[namespace].Get(id)
	if !has {
		return nil, auth.ErrNoAccount
	}
	return acct, nil
}

func (i *InMemoryAccountStore) FindByCredentials(ctx context.Context, namespace auth.AccountNamespace, creds map[string]string) (*auth.Account, error) {
	if _, hasNs := i.data[namespace]; !hasNs {
		return nil, fmt.Errorf("invalid namespace: %s", namespace)
	}
	for _, acct := range i.data[namespace].Items() {
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
	if _, hasNs := i.data[account.Namespace]; !hasNs {
		return fmt.Errorf("invalid namespace: %s", account.Namespace)
	}
	existing, _ := i.FindByCredentials(ctx, account.Namespace, creds)
	if existing != nil {
		return auth.ErrAccountExists
	}

	acctId := uuid.NewString()
	account.ID = acctId
	i.data[account.Namespace].Set(acctId, account)
	return nil
}

func (i *InMemoryAccountStore) Update(ctx context.Context, account *auth.Account) error {
	if _, hasNs := i.data[account.Namespace]; !hasNs {
		return fmt.Errorf("invalid namespace: %s", account.Namespace)
	}
	i.data[account.Namespace].Set(account.ID, account)
	return nil
}
