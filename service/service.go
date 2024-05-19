package service

import (
	"errors"
	"math/rand"

	"github.com/sirupsen/logrus"
)

type Service struct {
	db   *DB
	mail *EmailService
}

type EmailService struct{}

func NewService(db *DB) *Service {
	return &Service{db: db, mail: &EmailService{}}
}

func (s *Service) SendEmails(mailingDetails []*MailingDetails) ([]int, error) {
	IDsSent := make([]int, 0)
	for _, details := range mailingDetails {
		err := s.mail.SendEmail(details)
		if err != nil {
			logrus.WithError(err).WithField("details", details).Warn("failed to send email")
			continue
		}
		IDsSent = append(IDsSent, details.ID)
	}
	return IDsSent, nil
}

// mocked random behaviour 1/10 chance of failure
func (es *EmailService) SendEmail(mailingDetails *MailingDetails) error {
	if rand.Intn(10) == 0 {
		return errors.New("oops")
	}
	return nil
}
