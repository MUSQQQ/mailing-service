package storage

import (
	"mailing-service/cmd/mailing-service/config"

	"github.com/go-pg/pg/v10"
)

type DB struct {
	db *pg.DB
}

func New(cfg *config.DBConfig) *DB {
	pgDB := pg.Connect(&pg.Options{
		User:     cfg.User,
		Password: cfg.Password,
		Addr:     cfg.Address,
	})
	return &DB{
		db: pgDB,
	}
}

func (db *DB) Close() error {
	return db.db.Close()
}

//func (db *DB) InsertMailingDetails(){}
//func (db *DB) GetMailingDetails(){}
//func (db *DB) DeleteMailingDetails(){}
