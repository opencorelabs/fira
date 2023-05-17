package application

import (
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/auth/stores/in_memory"
	"github.com/opencorelabs/fira/internal/developer"
	"github.com/opencorelabs/fira/internal/developer/stores/app_memory"
)

func (a *App) AccountStore() auth.AccountStore {
	a.initMtx.Lock()
	defer a.initMtx.Unlock()

	if a.accountStore == nil {
		a.accountStore = in_memory.New()
	}

	return a.accountStore
}

func (a *App) AppStore() developer.AppStore {
	a.initMtx.Lock()
	defer a.initMtx.Unlock()

	if a.appStore == nil {
		a.appStore = app_memory.New()
	}

	return a.appStore
}
