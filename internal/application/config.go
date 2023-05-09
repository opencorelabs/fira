package application

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug       bool   `default:"true"`
	BindHTTP    string `default:"localhost:8080"`
	ClientDir   string `default:"./client"`
	FrontendURL string `default:"http://localhost:3000"`
}

func InitConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("fira", &cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to read environment: %w", err)
	}
	return &cfg, nil
}
