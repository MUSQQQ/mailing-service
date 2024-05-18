package service

import "time"

type MailingDetails struct {
	ID         int
	MailingID  int
	Email      string
	Title      string
	Content    string
	InsertTime time.Time
}
