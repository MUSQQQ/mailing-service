package service

import (
	"mailing-service/cmd/mailing-service/config"

	"github.com/go-pg/pg/v10"
)

type DB struct {
	db *pg.DB
}

func NewDB(cfg *config.DBConfig) *DB {
	pgDB := pg.Connect(&pg.Options{
		User:     cfg.User,
		Password: cfg.Password,
		Addr:     cfg.Address,
		Database: cfg.Database,
	})
	return &DB{
		db: pgDB,
	}
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) CreateMailingDetails(mailingDetails MailingDetails) error {
	_, err := db.db.Exec(
		`INSERT INTO mailing_details (email, title, content, mailing_id, insert_time) VALUES (?,?,?,?, ?)`,
		mailingDetails.Email,
		mailingDetails.Title,
		mailingDetails.Content,
		mailingDetails.MailingID,
		mailingDetails.InsertTime)
	return err
}

//func (db *DB) GetMailingDetails(){}
//func (db *DB) DeleteMailingDetails(){}
