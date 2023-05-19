package account_psql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/persistence/psql"
	"github.com/opencorelabs/fira/internal/persistence/snowflake"

	"go.uber.org/zap"
)

type PostgresStore struct {
	logger     *zap.Logger
	dbProvider psql.Provider
	sfProvider snowflake.Provider
}

type QueryRunner interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

func New(sfProvider snowflake.Provider, dbProvider psql.Provider) auth.AccountStore {
	return &PostgresStore{
		dbProvider: dbProvider,
		sfProvider: sfProvider,
	}
}

func (p *PostgresStore) FindAccountByID(ctx context.Context, namespace auth.AccountNamespace, id string) (*auth.Account, error) {
	sql := `SELECT * FROM accounts WHERE namespace = $1 AND account_id = $2`
	rows, queryErr := p.dbProvider.DB().Query(ctx, sql, namespace, id)
	if queryErr != nil {
		return nil, fmt.Errorf("error querying account: %w", queryErr)
	}
	return scanOneAccount(rows)
}

func (p *PostgresStore) FindByCredentials(ctx context.Context, namespace auth.AccountNamespace, creds map[string]string) (*auth.Account, error) {
	return p.findByCredentials(ctx, p.dbProvider.DB(), namespace, creds)
}

func (p *PostgresStore) findByCredentials(ctx context.Context, tx psql.DB, namespace auth.AccountNamespace, creds map[string]string) (*auth.Account, error) {
	sql := `SELECT * FROM accounts WHERE namespace = $1 AND credentials @> $2`
	rows, queryErr := tx.Query(ctx, sql, namespace, creds)
	if queryErr != nil {
		return nil, fmt.Errorf("error querying account: %w", queryErr)
	}
	return scanOneAccount(rows)
}

func (p *PostgresStore) Create(ctx context.Context, account *auth.Account, creds map[string]string) error {
	return pgx.BeginFunc(ctx, p.dbProvider.DB(), func(tx pgx.Tx) error {
		existing, _ := p.findByCredentials(ctx, tx, account.Namespace, creds)
		if existing != nil {
			return auth.ErrAccountExists
		}

		account.ID = p.sfProvider.Generator().Generate().String()

		if account.Credentials == nil {
			account.Credentials = make(map[string]string)
		}

		sql := `INSERT INTO accounts (
                	account_id, 
                	namespace, 
                	valid, 
                	credentials_type, 
                	credentials, 
                	name, 
					avatar_url, 
					email, 
					created_at, 
					updated_at
				) VALUES (
					$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
				)`
		_, err := tx.Exec(ctx, sql,
			account.ID,
			account.Namespace,
			account.Valid,
			account.CredentialsType,
			account.Credentials,
			account.Name,
			account.AvatarURL,
			account.Email,
			account.CreatedAt,
			account.UpdatedAt,
		)
		if err != nil {
			return fmt.Errorf("error creating account: %w", err)
		}
		return nil
	})
}

func (p *PostgresStore) Update(ctx context.Context, account *auth.Account) error {
	sql := `
		UPDATE accounts SET
			valid = $1,
			credentials_type = $2,
			credentials = $3,
			name = $4,
			avatar_url = $5,
			email = $6,
			updated_at = $7
		WHERE namespace = $8 AND account_id = $9`
	_, err := p.dbProvider.DB().Exec(ctx, sql,
		account.Valid,
		account.CredentialsType,
		account.Credentials,
		account.Name,
		account.AvatarURL,
		account.Email,
		account.UpdatedAt,
		account.Namespace,
		account.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating account: %w", err)
	}
	return nil
}

func scanOneAccount(rows pgx.Rows) (*auth.Account, error) {
	acct, scanErr := pgx.CollectOneRow(rows, pgx.RowToStructByName[auth.Account])
	if scanErr != nil {
		if scanErr == pgx.ErrNoRows {
			return nil, auth.ErrNoAccount
		} else {
			return nil, fmt.Errorf("error scanning account: %w", scanErr)
		}
	}
	return &acct, nil
}
