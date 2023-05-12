package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Provider interface {
	Config() *Config
}

type Config struct {
	Debug       bool   `default:"true"`
	BindHttp    string `default:"localhost:8080" split_words:"true"`
	ClientDir   string `default:"./client" split_words:"true"`
	FrontendUrl string `default:"http://localhost:3000" split_words:"true"`
	GrpcUrl     string `default:"localhost:5567"`
}

func Init() (*Config, error) {
	var cfg Config
	err := envconfig.Process("fira", &cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to read environment: %w", err)
	}
	return &cfg, nil
}

func (c *Config) Config() *Config {
	return c
}
