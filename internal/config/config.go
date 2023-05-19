package config

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"
)

type Provider interface {
	Config() *Config
}

type Config struct {
	Debug         bool   `default:"true"`
	BindHttp      string `default:"localhost:8080" split_words:"true"`
	ClientDir     string `default:"./client" split_words:"true"`
	FrontendUrl   string `default:"http://localhost:3000" split_words:"true"`
	GrpcUrl       string `default:"localhost:5567"`
	NodeId        int64  `default:"1" split_words:"true"`
	PostgresUrl   string `default:"postgres://postgres:docker@localhost:5432/fira?sslmode=disable" split_words:"true"`
	MigrationsDir string `default:"./pg/migrations" split_words:"true"`
	pgpoolConfig  *pgxpool.Config
}

func Init() (*Config, error) {
	var cfg Config
	err := envconfig.Process("fira", &cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to read environment: %w", err)
	}
	pgpoolConfig, pgpoolErr := pgxpool.ParseConfig(cfg.PostgresUrl)
	if pgpoolErr != nil {
		return nil, fmt.Errorf("unable to parse postgres url: %w", pgpoolErr)
	}
	cfg.pgpoolConfig = pgpoolConfig
	return &cfg, nil
}

func (c *Config) Config() *Config {
	return c
}

func (c *Config) PgxPoolConfig() *pgxpool.Config {
	return c.pgpoolConfig
}
