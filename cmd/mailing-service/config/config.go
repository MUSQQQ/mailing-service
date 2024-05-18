package config

import (
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	DBConfig     *DBConfig
	RouterConfig *RouterConfig
}

type DBConfig struct {
	User     string `env:"DB_USER" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD" envDefault:"postgres"`
	Address  string `env:"DB_ADDRESS" envDefault:"postgres:5432"`
	Database string `env:"DB_NAME" envDefault:"mailing"`
}

type RouterConfig struct {
	Port string `env:"PORT" envDefault:"8080"`
}

func Parse() *Config {
	dbCfg := &DBConfig{}
	if err := env.Parse(dbCfg); err != nil {
		log.Panicf("unable to parse db config: %v", err)
	}
	routerCfg := &RouterConfig{}
	if err := env.Parse(routerCfg); err != nil {
		log.Panicf("unable to parse router config: %v", err)
	}
	return &Config{
		DBConfig:     dbCfg,
		RouterConfig: routerCfg,
	}
}
