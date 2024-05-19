package main

import (
	"context"
	"mailing-service/cmd/mailing-service/config"
	"mailing-service/service"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("starting mailing-service")
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	cfg := config.Parse()

	db := service.NewDB(cfg.DBConfig)

	srv := service.NewService(db)

	router := service.NewRouter(cfg.RouterConfig)

	router.RegisterHandlers(srv)
	go router.Run(ctx)

	cron := service.NewCron(cfg.CronConfig, db)
	go cron.Run(ctx)
	<-ctx.Done()
}
