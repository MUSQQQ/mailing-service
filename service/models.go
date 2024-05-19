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

type CreateMailingDetailsRequest struct {
	Email      string `json:"email"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	MailingID  int    `json:"mailing_id"`
	InsertTime string `json:"insert_time"`
}

func (req *CreateMailingDetailsRequest) toMailingDetails() (*MailingDetails, error) {
	time, err := time.Parse(time.RFC3339, req.InsertTime)
	if err != nil {
		return nil, err
	}
	return &MailingDetails{
		Email:      req.Email,
		Title:      req.Title,
		Content:    req.Content,
		MailingID:  req.MailingID,
		InsertTime: time,
	}, nil
}

type SendMailingRequest struct {
	MailingID int `json:"mailing_id"`
}
