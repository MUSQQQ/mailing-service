package service

import (
	"fmt"
	"mailing-service/cmd/mailing-service/config"
	"strconv"
	"strings"
	"time"

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
		`INSERT INTO mailing_details (email, title, content, mailing_id, insert_time) VALUES (?,?,?,?,?)`,
		mailingDetails.Email,
		mailingDetails.Title,
		mailingDetails.Content,
		mailingDetails.MailingID,
		mailingDetails.InsertTime)
	return err
}

func (db *DB) GetMailingDetailsByMailingID(mailingID int) ([]*MailingDetails, error) {
	var details []*MailingDetails

	_, err := db.db.Query(
		&details,
		`SELECT id, mailing_id, email, title, content, insert_time FROM mailing_details WHERE mailing_id=?`,
		mailingID,
	)
	if err != nil {
		return nil, err
	}
	return details, nil
}

func (db *DB) DeleteMailingDetails(whereClauses []string) error {
	_, err := db.db.Exec(fmt.Sprintf(`DELETE FROM mailing_details WHERE %s`, strings.Join(whereClauses, " AND ")))

	return err
}

type Filters struct {
	clauses []string
}

func (f *Filters) ByIDs(ids []int) *Filters {
	if f.clauses == nil {
		f.clauses = []string{}
	}
	idsStr := []string{}
	for _, id := range ids {
		idsStr = append(idsStr, strconv.Itoa(id))
	}
	f.clauses = append(f.clauses, fmt.Sprintf(`id IN (%v)`, strings.Join(idsStr, ", ")))
	return f
}

func (f *Filters) ByInsertTimeBefore(insertTime time.Time) *Filters {
	if f.clauses == nil {
		f.clauses = []string{}
	}
	f.clauses = append(f.clauses, fmt.Sprintf(`insert_time<'%s'`, insertTime.UTC().Format("2006-01-02 15:04:05")))
	return f
}

func (f *Filters) ByMailingID(id int) *Filters {
	if f.clauses == nil {
		f.clauses = []string{}
	}
	f.clauses = append(f.clauses, fmt.Sprintf(`mailing_id=%d`, id))
	return f
}
