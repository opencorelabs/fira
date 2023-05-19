package application

import (
	"context"
	snowflakelib "github.com/bwmarrin/snowflake"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/opencorelabs/fira/internal/persistence/psql"
	"github.com/opencorelabs/fira/internal/persistence/snowflake"
	"go.uber.org/zap"
)

func (a *App) StartDB(ctx context.Context) {
	db := a.DB()
	a.StartService(ctx, "db-pool", func(ctx context.Context, errChan chan error) Finalizer {
		return func(ctx context.Context) error {
			db.Close()
			return nil
		}
	})
}

func (a *App) DB() psql.DBCloser {
	a.initMtx.Lock()
	defer a.initMtx.Unlock()

	if a.pgxPool == nil {
		var newErr error
		a.pgxPool, newErr = pgxpool.NewWithConfig(context.Background(), a.cfg.PgxPoolConfig())
		if newErr != nil {
			a.logger.Fatal("error initializing pgx pool", zap.Error(newErr))
		}
	}

	return a.pgxPool
}

func (a *App) Generator() snowflake.Generator {
	a.initMtx.Lock()
	defer a.initMtx.Unlock()

	if a.sfNode == nil {
		var newErr error
		a.sfNode, newErr = snowflakelib.NewNode(a.cfg.NodeId)
		if newErr != nil {
			a.logger.Fatal("error initializing snowflake generator", zap.Error(newErr))
		}
	}

	return a.sfNode
}
