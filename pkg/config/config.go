package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	PGAddress  string `env:"PG_ADDRESS" envDefault:"host.docker.internal:5432"`
	PGUser     string `env:"PG_USER" envDefault:"beer"`
	PGPassword string `env:"PG_PASSWORD" envDefault:"beer"`
	PGDatabase string `env:"PG_DATABASE" envDefault:"beer"`

	ServerPort string `env:"SERVER_PORT" envDefault:"3300"`
}

func GetConfig() (*Config, error) {
	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
