package service

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s *Service) CreateMailingDetails(w http.ResponseWriter, r *http.Request) {
	logrus.Info("start handling CreateMailingDetails")
	req := &CreateMailingDetailsRequest{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logrus.WithError(err).Error("failed to decode request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mailingDetails, err := req.toMailingDetails()
	if err != nil {
		logrus.WithError(err).Error("failed to map request to mailing details")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	logrus.WithField("details", mailingDetails).Info("mapped")

	err = s.db.CreateMailingDetails(*mailingDetails)
	if err != nil {
		logrus.WithError(err).Error("failed to save mailing details")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Service) SendMailing(w http.ResponseWriter, r *http.Request) {
	logrus.Info("start handling SendMailingRequest")
	req := &SendMailingRequest{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logrus.WithError(err).Error("failed to decode request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	details, err := s.db.GetMailingDetailsByMailingID(req.MailingID)
	if err != nil {
		logrus.WithError(err).Error("failed to get mailing details")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logrus.WithField("details", len(details)).Info("got details to send")

	IDsSent, err := s.SendEmails(details)
	if err != nil {
		logrus.WithError(err).Error("failed to send emails")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//deleting that what was sent

	if len(details) > len(IDsSent) {
		info, _ := json.Marshal(map[string]string{
			"info": "partial_success",
		})
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(info)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Service) DeleteMailingDetails(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
