package application

import (
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/auth/stores/account_psql"
	"github.com/opencorelabs/fira/internal/developer"
	"github.com/opencorelabs/fira/internal/developer/stores/app_psql"
)

func (a *App) AccountStore() auth.AccountStore {
	a.initMtx.Lock()
	defer a.initMtx.Unlock()

	if a.accountStore == nil {
		a.accountStore = account_psql.New(a, a)
	}

	return a.accountStore
}

func (a *App) AppStore() developer.AppStore {
	a.initMtx.Lock()
	defer a.initMtx.Unlock()

	if a.appStore == nil {
		a.appStore = app_psql.New(a, a)
	}

	return a.appStore
}
