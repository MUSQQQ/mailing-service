package main

import (
	"mailing-service/cmd/mailing-service/config"
	"mailing-service/service"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("starting mailing-service")

	cfg := config.Parse()

	db := service.NewDB(cfg.DBConfig)

	srv := service.New(db)

	router := service.NewRouter(cfg.RouterConfig)

	router.RegisterHandlers(srv)
	router.Run()

}
