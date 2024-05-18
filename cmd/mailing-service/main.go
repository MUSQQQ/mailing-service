package main

import (
	"mailing-service/cmd/mailing-service/config"
	"mailing-service/service"
	"mailing-service/storage"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("starting mailing-service")

	cfg := config.Parse()

	db := storage.New(cfg.DBConfig)

	srv := service.New(db)

	router := service.NewRouter(cfg.RouterConfig)

	router.RegisterHandlers(srv)
	router.Run()

}
