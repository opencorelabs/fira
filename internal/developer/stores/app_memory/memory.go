package app_memory

import (
	"context"
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

func (s *Store) CreateApp(ctx context.Context, app *developer.App) error {
	app.ID = uuid.NewString()
	s.apps.Set(app.ID, app)
	return nil
}

func (s *Store) UpdateApp(ctx context.Context, app *developer.App) error {
	s.apps.Set(app.ID, app)
	return nil
}

func (s *Store) GetAppsByAccountID(ctx context.Context, accountID string) (apps []*developer.App, err error) {
	for _, v := range s.apps.Items() {
		if v.AccountID == accountID {
			apps = append(apps, v)
		}
	}
	return
}

func (s *Store) GetAppByID(ctx context.Context, appID string) (*developer.App, error) {
	app, exists := s.apps.Get(appID)
	if !exists {
		return nil, fmt.Errorf("app not found: %s", appID)
	}
	return app, nil
}
