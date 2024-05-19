package config

import (
	"log"
	"time"

	"github.com/caarlos0/env"
)

type Config struct {
	DBConfig     *DBConfig
	RouterConfig *RouterConfig
	CronConfig   *CronConfig
}

type DBConfig struct {
	User     string `env:"DB_USER" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD" envDefault:"postgres"`
	Host     string `env:"DB_HOST" envDefault:"postgres"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	Database string `env:"DB_NAME" envDefault:"mailing"`
}

type RouterConfig struct {
	Port string `env:"PORT" envDefault:"8080"`
}

type CronConfig struct {
	Interval time.Duration `env:"CRON_INTERVAL" envDefault:"60s"`
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
	cronCfg := &CronConfig{}
	if err := env.Parse(cronCfg); err != nil {
		log.Panicf("unable to parse router config: %v", err)
	}
	return &Config{
		DBConfig:     dbCfg,
		RouterConfig: routerCfg,
		CronConfig:   cronCfg,
	}
}
