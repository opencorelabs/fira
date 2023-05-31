package application

import (
	"github.com/opencorelabs/fira/internal/backend"
	"github.com/opencorelabs/fira/internal/backend/aggregators/plaid"
)

func (a *App) Backend() backend.Interface {
	a.initMtx.Lock()
	defer a.initMtx.Unlock()

	if a.backend == nil {
		a.backend = backend.NewMetaInterface(a, []backend.Interface{
			plaid.New(a, a.cfg.Plaid.ClientId, a.cfg.Plaid.ClientSecret, a.cfg.Plaid.Environment),
		})
	}

	return a.backend
}
