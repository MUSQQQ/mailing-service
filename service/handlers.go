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
	logrus.WithField("body", req).Info("unmarshalled")

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
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *Service) DeleteMailingDetails(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
