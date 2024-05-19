package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func (s *Service) CreateMailingDetails(w http.ResponseWriter, r *http.Request) {
	log := logrus.NewEntry(logrus.New())
	log.Info("handling CreateMailingDetails")
	req := &CreateMailingDetailsRequest{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.WithError(err).Error("failed to decode request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log = log.WithField("request_body", req)

	mailingDetails, err := req.toMailingDetails()
	if err != nil {
		log.WithError(err).Error("failed to map request to mailing details")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.WithField("mapped_details", mailingDetails)

	err = s.db.CreateMailingDetails(mailingDetails)
	if err != nil {
		log.WithError(err).Error("failed to save mailing details")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Info("saved mailing details")
	w.WriteHeader(http.StatusOK)
}

func (s *Service) SendMailing(w http.ResponseWriter, r *http.Request) {
	log := logrus.NewEntry(logrus.New())
	log.Info("handling SendMailing")
	req := &SendMailingRequest{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.WithError(err).Error("failed to decode request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log = log.WithField("request_body", req)

	details, err := s.db.GetMailingDetailsByMailingID(req.MailingID)
	if err != nil {
		log.WithError(err).Error("failed to get mailing details")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log = log.WithField("mailing_details_found", len(details))

	IDsSent, err := s.SendEmails(details)
	if err != nil {
		log.WithError(err).Error("failed to send emails")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log = log.WithField("emails_sent", len(IDsSent))

	filters := &Filters{
		clauses: []string{},
	}
	err = s.db.DeleteMailingDetails(filters.ByIDs(IDsSent).clauses)
	if err != nil {
		log.WithError(err).Error("failed to delete mailing details")
	}

	if len(details) > len(IDsSent) {
		log.Warn("failed to send some emails")
		info, _ := json.Marshal(map[string]string{
			"info": "partial_success",
		})
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(info)
		return
	}

	log.Info("all emails successfully sent")
	w.WriteHeader(http.StatusOK)
}

func (s *Service) DeleteMailingDetails(w http.ResponseWriter, r *http.Request) {
	log := logrus.NewEntry(logrus.New())
	log.Info("handling DeleteMailingDetails")

	mailingID := chi.URLParam(r, "mailingID")
	if mailingID == "" {
		log.Warn("mailingID not found")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log = log.WithField("mailing_id", mailingID)

	id, err := strconv.Atoi(mailingID)
	if err != nil {
		log.Warn("incorrect mailingID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	filters := Filters{clauses: []string{}}
	err = s.db.DeleteMailingDetails(filters.ByMailingID(id).clauses)
	if err != nil {
		log.Warn("failed to delete mailing details")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info("mailing details deleted")
	w.WriteHeader(http.StatusOK)
}
