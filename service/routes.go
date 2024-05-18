package service

import (
	"mailing-service/cmd/mailing-service/config"
	"net/http"

	"github.com/go-chi/chi"
)

type Router struct {
	router *chi.Mux
	Port   string
}

func NewRouter(cfg *config.RouterConfig) *Router {
	return &Router{
		router: chi.NewRouter(),
		Port:   cfg.Port,
	}
}

func (r *Router) RegisterHandlers(service *Service) {
	r.router.Route("/api/messages", func(r chi.Router) {
		r.Post("/", service.CreateMailingDetails)
		r.Post("/send", service.SendMailing)
		r.Delete("/{mailingID}", service.DeleteMailingDetails)
	})
}

func (r *Router) Run() {
	http.ListenAndServe(":"+r.Port, r.router)
}
