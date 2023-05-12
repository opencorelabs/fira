package application

import (
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/auth/stores/in_memory"
)

func (a *App) AccountStore() auth.AccountStore {
	a.initMtx.Lock()
	defer a.initMtx.Unlock()

	if a.accountStore == nil {
		a.accountStore = in_memory.New()
	}

	return a.accountStore
}
