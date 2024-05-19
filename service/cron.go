package service

import (
	"context"
	"mailing-service/cmd/mailing-service/config"
	"time"

	"github.com/sirupsen/logrus"
)

type Cron struct {
	interval time.Duration
	db       *DB
}

func NewCron(cfg *config.CronConfig, db *DB) *Cron {
	return &Cron{
		interval: cfg.Interval,
		db:       db,
	}
}

func (c *Cron) Run(ctx context.Context) {
	ticker := time.NewTicker(c.interval)
	for {
		select {
		case <-ticker.C:
			err := c.RemoveMailingDetails()
			if err != nil {
				logrus.WithError(err).Error("failed to delete mailing details older than 5 minutes")
			}
		case <-ctx.Done():
			logrus.Info("cron shutdown")
			return
		}
	}
}

func (c *Cron) RemoveMailingDetails() error {
	filters := &Filters{clauses: []string{}}
	err := c.db.DeleteMailingDetails(filters.ByInsertTimeBefore(time.Now().Add(-5 * time.Minute).UTC()).clauses)
	return err
}
