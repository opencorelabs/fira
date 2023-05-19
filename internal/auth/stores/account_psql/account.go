package account_psql

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/opencorelabs/fira/internal/auth"
	"go.uber.org/zap"
)

type PostgresStore struct {
	logger    *zap.Logger
	pool      *pgxpool.Pool
	snowflake *snowflake.Node
}

func New() auth.AccountStore {
	generator, snowflakeErr := snowflake.NewNode(1)
	if snowflakeErr != nil {
		panic(snowflakeErr)
	}
	cfg, cfgErr := pgxpool.ParseConfig("postgres://postgres:docker@localhost:5432/fira?sslmode=disable")
	if cfgErr != nil {
		panic(cfgErr)
	}
	pool, poolErr := pgxpool.NewWithConfig(context.Background(), cfg)
	if poolErr != nil {
		panic(poolErr)
	}
	return &PostgresStore{
		logger:    zap.L().Named("account_psql"),
		pool:      pool,
		snowflake: generator,
	}
}

func (p *PostgresStore) FindAccountByID(ctx context.Context, namespace auth.AccountNamespace, id string) (*auth.Account, error) {
	sql := `
		SELECT ` + p.accountColumns() + `
		FROM accounts WHERE namespace = $1 AND account_id = $2
	`
	row := p.pool.QueryRow(ctx, sql, namespace, id)
	acct := &auth.Account{}
	err := p.scanColumns(acct, row)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, auth.ErrNoAccount
		} else {
			return nil, fmt.Errorf("error scanning account: %w", err)
		}
	}
	return acct, nil
}

func (p *PostgresStore) FindByCredentials(ctx context.Context, namespace auth.AccountNamespace, creds map[string]string) (*auth.Account, error) {
	var acct *auth.Account
	err := pgx.BeginFunc(ctx, p.pool, func(tx pgx.Tx) error {
		var findErr error
		acct, findErr = p.findByCredentials(ctx, tx, namespace, creds)
		return findErr
	})
	return acct, err
}

func (p *PostgresStore) findByCredentials(ctx context.Context, tx pgx.Tx, namespace auth.AccountNamespace, creds map[string]string) (*auth.Account, error) {
	sql := `
	SELECT ` + p.accountColumns() + `
	FROM accounts WHERE namespace = $1 AND credentials @> $2`
	row := tx.QueryRow(ctx, sql, namespace, creds)
	acct := &auth.Account{}
	err := p.scanColumns(acct, row)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, auth.ErrNoAccount
		} else {
			return nil, fmt.Errorf("error scanning account: %w", err)
		}
	}
	return acct, nil
}

func (p *PostgresStore) Create(ctx context.Context, account *auth.Account, creds map[string]string) error {
	return pgx.BeginFunc(ctx, p.pool, func(tx pgx.Tx) error {
		existing, _ := p.findByCredentials(ctx, tx, account.Namespace, creds)
		if existing != nil {
			return auth.ErrAccountExists
		}

		account.ID = p.snowflake.Generate().String()

		if account.Credentials == nil {
			account.Credentials = make(map[string]string)
		}

		sql := `INSERT INTO accounts (` + p.accountColumns() + `) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
		_, err := p.pool.Exec(ctx, sql, p.saveColumns(account)...)
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
	args := p.updateColumns(account)
	args = append(args, account.Namespace, account.ID)
	_, err := p.pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("error updating account: %w", err)
	}
	return nil
}

func (p *PostgresStore) accountColumns() string {
	return "account_id, namespace, valid, credentials_type, credentials, name, avatar_url, email, created_at, updated_at"
}

func (p *PostgresStore) scanColumns(a *auth.Account, row pgx.Row) error {
	return row.Scan(
		&a.ID,
		&a.Namespace,
		&a.Valid,
		&a.CredentialsType,
		&a.Credentials,
		&a.Name,
		&a.AvatarURL,
		&a.Email,
		&a.CreatedAt,
		&a.UpdatedAt,
	)
}

func (p *PostgresStore) saveColumns(a *auth.Account) []interface{} {
	return []interface{}{
		a.ID,
		a.Namespace,
		a.Valid,
		a.CredentialsType,
		a.Credentials,
		a.Name,
		a.AvatarURL,
		a.Email,
		a.CreatedAt,
		a.UpdatedAt,
	}
}

func (p *PostgresStore) updateColumns(a *auth.Account) []interface{} {
	return []interface{}{
		a.Valid,
		a.CredentialsType,
		a.Credentials,
		a.Name,
		a.AvatarURL,
		a.Email,
		a.UpdatedAt,
	}
}
