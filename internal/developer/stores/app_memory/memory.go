package app_memory

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/opencorelabs/fira/internal/developer"
	cmap "github.com/orcaman/concurrent-map/v2"
)

type Store struct {
	apps cmap.ConcurrentMap[string, *developer.App]
}

func New() developer.AppStore {
	return &Store{
		apps: cmap.New[*developer.App](),
	}
}

func (s *Store) CreateApp(app *developer.App) error {
	app.ID = uuid.NewString()
	s.apps.Set(app.ID, app)
	return nil
}

func (s *Store) GetAppsByAccountID(accountID string) (apps []*developer.App, err error) {
	for _, v := range s.apps.Items() {
		if v.AccountID == accountID {
			apps = append(apps, v)
		}
	}
	return
}

func (s *Store) GetAppByID(appID string) (*developer.App, error) {
	app, exists := s.apps.Get(appID)
	if !exists {
		return nil, fmt.Errorf("app not found: %s", appID)
	}
	return app, nil
}
