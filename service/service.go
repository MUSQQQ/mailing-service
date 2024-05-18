package service

import (
	"mailing-service/storage"
	"net/http"
)

type Service struct {
	db *storage.DB
}

func New(db *storage.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateMailingDetails(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte(`{"message": "Not Found"}`))
}

func (s *Service) SendMailing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte(`{"message": "Not Found"}`))
}

func (s *Service) DeleteMailingDetails(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte(`{"message": "Not Found"}`))
}
