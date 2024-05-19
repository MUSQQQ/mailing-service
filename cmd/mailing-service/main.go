package main

import (
	"context"
	"mailing-service/cmd/mailing-service/config"
	"mailing-service/service"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.NewEntry(logrus.New())
	log.Info("starting mailing-service")
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	log.Info("parsing config")
	cfg := config.Parse()

	log.Info("initializing database")
	db, err := service.NewDB(cfg.DBConfig)
	if err != nil {
		log.WithError(err).Fatal("failed to initialize DB")
	}

	log.Info("initializing service")
	srv := service.NewService(db)

	log.Info("initializing router")
	router := service.NewRouter(cfg.RouterConfig)

	log.Info("running router")
	router.RegisterHandlers(srv)
	go router.Run(ctx)

	log.Info("initializing cron")
	cron := service.NewCron(cfg.CronConfig, db)
	log.Info("running cron")
	go cron.Run(ctx)
	<-ctx.Done()
}
