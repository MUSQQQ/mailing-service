package service

import (
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	DBConfig
}

type DBConfig struct {
	User     string `env:"DB_USER" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD" envDefault:"postgres"`
	Address  string `env:"DB_ADDRESS" envDefault:"a"`
}

func Parse() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Panicf("unable to parse config: %v", err)
	}
	return &cfg
}
