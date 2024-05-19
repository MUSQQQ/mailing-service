package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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

	filters := &Filters{
		clauses: []string{},
	}
	err = s.db.DeleteMailingDetails(filters.ByIDs(IDsSent).clauses)
	if err != nil {
		logrus.WithError(err).Error("failed to delete mailing details")
	}

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
	mailingID := chi.URLParam(r, "mailingID")
	if mailingID == "" {
		logrus.Warn("mailingID not found")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(mailingID)
	if err != nil {
		logrus.Warn("incorrect mailingID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	filters := Filters{clauses: []string{}}
	err = s.db.DeleteMailingDetails(filters.ByMailingID(id).clauses)
	if err != nil {
		logrus.Warn("failed to delete mailing details")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
