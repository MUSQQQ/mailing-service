package service

import (
	"context"
	"mailing-service/cmd/mailing-service/config"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type RouterInterface interface {
	RegisterHandlers(service *Service)
	Run(ctx context.Context)
}

type Router struct {
	Router *chi.Mux
	Port   string
}

var _ RouterInterface = (*Router)(nil)

func NewRouter(cfg *config.RouterConfig) *Router {
	return &Router{
		Router: chi.NewRouter(),
		Port:   cfg.Port,
	}
}

func (r *Router) RegisterHandlers(service *Service) {
	r.Router.Route("/api/messages", func(r chi.Router) {
		r.Post("/", service.CreateMailingDetails)
		r.Post("/send", service.SendMailing)
		r.Delete("/{mailingID}", service.DeleteMailingDetails)
	})
}

func (r *Router) Run(ctx context.Context) {
	srv := http.Server{
		Addr:    ":" + r.Port,
		Handler: r.Router,
	}
	go srv.ListenAndServe()

	<-ctx.Done()
	logrus.Info("router shutdown")
	srv.Shutdown(ctx)
}
