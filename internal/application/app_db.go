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
)

func (a *App) StartDB(ctx context.Context) {
	if a.cfg.LocalPostgres.Enable {
		a.StartService(ctx, "embedded-postgres", func(ctx context.Context, errChan chan error) Finalizer {
			defer a.PanicRecovery(errChan)

			logBuf := newBuf()
			postgres := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().
				Username(a.cfg.LocalPostgres.Username).
				Password(a.cfg.LocalPostgres.Password).
				Database(a.cfg.LocalPostgres.Database).
				Version(embeddedpostgres.V15).
				RuntimePath("/tmp/pg-runtime").
				DataPath(a.cfg.LocalPostgres.DataPath).
				BinariesPath(a.cfg.LocalPostgres.BinariesPath).
				Logger(logBuf),
			)

			logger := a.Logger().Named("embedded-postgres")

			go a.pipeToLogger(logBuf, zap.InfoLevel, logger, errChan)

			logger.Info("starting embedded postgres")

			errChan <- postgres.Start()

			return func(ctx context.Context) error {
				return postgres.Stop()
			}
		})
	}

	db := a.DB()
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
