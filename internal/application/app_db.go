package application

import (
	"bytes"
	"context"
	snowflakelib "github.com/bwmarrin/snowflake"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/opencorelabs/fira/internal/persistence/psql"
	"github.com/opencorelabs/fira/internal/persistence/snowflake"
	"go.uber.org/zap"
	"io"
	"sync"
	"time"
)

func (a *App) StartDB(ctx context.Context) {
	if a.cfg.EmbeddedPostgres.Enable {
		startChan := make(chan struct{})
		defer close(startChan)
		semOnce := sync.Once{}

		a.StartService(ctx, "embedded-postgres", func(ctx context.Context, errChan chan error) Finalizer {
			var postgres *embeddedpostgres.EmbeddedPostgres
			logger := a.Logger().Named("embedded-postgres")

			go func() {
				defer a.PanicRecovery(errChan)

				logBuf := newBuf()
				postgres = embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().
					Username(a.cfg.EmbeddedPostgres.Username).
					Password(a.cfg.EmbeddedPostgres.Password).
					Database(a.cfg.EmbeddedPostgres.Database).
					Port(uint32(a.cfg.EmbeddedPostgres.Port)).
					Version(embeddedpostgres.V15).
					RuntimePath("/tmp/pg-runtime").
					DataPath(a.cfg.EmbeddedPostgres.DataPath).
					BinariesPath(a.cfg.EmbeddedPostgres.BinariesPath).
					Logger(logBuf),
				)

				go a.pipeToLogger(logBuf, zap.InfoLevel, logger, errChan)

				logger.Info("starting embedded postgres")

				startErr := postgres.Start()

				if startErr == nil {
					semOnce.Do(func() {
						startChan <- struct{}{}
					})
				} else {
					errChan <- startErr
				}
			}()

			return func(ctx context.Context) error {
				if postgres != nil {
					logger.Info("stopping embedded postgres")
					return postgres.Stop()
				}
				return nil
			}
		})

		select {
		case <-time.After(15 * time.Second):
			a.Logger().Fatal("unable to start embedded postgres")
		case <-startChan:
			a.Logger().Info("embedded postgres started")
		}
	}

	db := a.DB()
	rows, qErr := db.Query(ctx, "SELECT 1")
	if qErr != nil {
		a.Logger().Fatal("unable to query db", zap.Error(qErr))
	}
	defer rows.Close()
	has := rows.Next()
	if !has {
		a.Logger().Fatal("unable to query db")
	}

	migrateErr := psql.Migrate(a.cfg.MigrationsDir, a.cfg.PostgresUrl)
	if migrateErr != nil {
		a.Logger().Fatal("unable to migrate database", zap.Error(migrateErr))
	}

	a.StartService(ctx, "db-pool", func(ctx context.Context, errChan chan error) Finalizer {
		return func(ctx context.Context) error {
			db.Close()
			return nil
		}
	})
}

type buf struct {
	buf *bytes.Buffer
}

func (b *buf) Read(p []byte) (n int, err error) {
	return b.buf.Read(p)
}

func (b *buf) Write(p []byte) (n int, err error) {
	return b.buf.Write(p)
}

func (b *buf) Close() error {
	return nil
}

func newBuf() interface {
	io.ReadWriter
	io.Closer
} {
	return &buf{
		buf: &bytes.Buffer{},
	}
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
