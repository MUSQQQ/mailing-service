package service

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewMockDB(db *sql.DB) *DB {
	return &DB{
		db: db,
	}
}

func TestCreateMailingDetailsDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed creating sqlmock: %v", err)
	}
	defer db.Close()

	testDB := NewMockDB(db)

	tNow := time.Now()

	md := &MailingDetails{
		Email:      "test@test.com",
		Title:      "title",
		Content:    "content",
		MailingID:  1,
		InsertTime: tNow,
	}

	mock.ExpectExec("INSERT INTO mailing_details (.+) VALUES").
		WithArgs(md.Email, md.Title, md.Content, md.MailingID, md.InsertTime).
		WillReturnResult(sqlmock.NewResult(int64(1), int64(1)))

	err = testDB.CreateMailingDetails(md)
	if err != nil {
		t.Errorf("unexpected sql error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations not met: %v", err)
	}
}
