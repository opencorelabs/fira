package app_psql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/opencorelabs/fira/internal/developer"
	"github.com/opencorelabs/fira/internal/persistence/psql"
	"github.com/opencorelabs/fira/internal/persistence/snowflake"
	"time"
)

type PostgresStore struct {
	snowflakeProvider snowflake.Provider
	dbProvider        psql.Provider
}

func New(snowflakeProvider snowflake.Provider, dbProvider psql.Provider) developer.AppStore {
	return &PostgresStore{
		snowflakeProvider: snowflakeProvider,
		dbProvider:        dbProvider,
	}
}

func (p *PostgresStore) CreateApp(ctx context.Context, app *developer.App) error {
	app.ID = p.snowflakeProvider.Generator().Generate().String()
	app.CreatedAt = time.Now()
	app.UpdatedAt = time.Now()
	if app.Tokens == nil {
		app.Tokens = make(developer.TokenMap)
	}

	sql := `
		INSERT INTO apps (
			app_id, 
			name, 
			account_id, 
			tokens,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := p.dbProvider.DB().Exec(ctx, sql,
		app.ID,
		app.Name,
		app.AccountID,
		app.Tokens,
		app.CreatedAt,
		app.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create app: %w", err)
	}
	return nil
}

func (p *PostgresStore) UpdateApp(ctx context.Context, app *developer.App) error {
	app.UpdatedAt = time.Now()
	sql := `
		UPDATE apps SET 
			name = $1,
			tokens = $2,
			updated_at = $3
		WHERE app_id = $4
	`
	_, err := p.dbProvider.DB().Exec(ctx, sql,
		app.Name,
		app.Tokens,
		app.UpdatedAt,
		app.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update app: %w", err)
	}
	return nil
}

func (p *PostgresStore) GetAppsByAccountID(ctx context.Context, accountID string) ([]*developer.App, error) {
	sql := `SELECT * FROM apps WHERE account_id = $1 ORDER BY created_at DESC`
	rows, err := p.dbProvider.DB().Query(ctx, sql, accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to get apps by account id: %w", err)
	}
	defer rows.Close()
	apps, scanErr := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[developer.App])
	if scanErr != nil {
		return nil, fmt.Errorf("failed to scan apps: %w", scanErr)
	}
	return apps, nil
}

func (p *PostgresStore) GetAppByID(ctx context.Context, appID string) (*developer.App, error) {
	sql := `SELECT * FROM apps WHERE app_id = $1`
	rows, err := p.dbProvider.DB().Query(ctx, sql, appID)
	if err != nil {
		return nil, fmt.Errorf("failed to get app by id: %w", err)
	}
	defer rows.Close()
	app, scanErr := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[developer.App])
	if scanErr != nil {
		return nil, fmt.Errorf("failed to scan app: %w", scanErr)
	}
	return app, nil
}
