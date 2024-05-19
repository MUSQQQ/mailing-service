package service

import (
	"database/sql"
	"fmt"
	"mailing-service/cmd/mailing-service/config"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . DBInterface

type DBInterface interface {
	CreateMailingDetails(mailingDetails *MailingDetails) error
	GetMailingDetailsByMailingID(mailingID int) ([]*MailingDetails, error)
	DeleteMailingDetails(whereClauses []string) error
}

type DB struct {
	db *sql.DB
}

var _ DBInterface = (*DB)(nil)

func NewDB(cfg *config.DBConfig) (*DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database))
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) CreateMailingDetails(mailingDetails *MailingDetails) error {
	_, err := db.db.Exec(
		`INSERT INTO mailing_details (email, title, content, mailing_id, insert_time) VALUES ($1,$2,$3,$4,$5)`,
		mailingDetails.Email,
		mailingDetails.Title,
		mailingDetails.Content,
		mailingDetails.MailingID,
		mailingDetails.InsertTime)
	return err
}

func (db *DB) GetMailingDetailsByMailingID(mailingID int) ([]*MailingDetails, error) {
	var details []*MailingDetails

	rows, err := db.db.Query(
		`SELECT id, mailing_id, email, title, content, insert_time FROM mailing_details WHERE mailing_id=$1`,
		mailingID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		detail := &MailingDetails{}
		rows.Scan(
			&detail.ID,
			&detail.MailingID,
			&detail.Email,
			&detail.Title,
			&detail.Content,
			&detail.InsertTime,
		)
		details = append(details, detail)
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
