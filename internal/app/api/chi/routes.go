package chi

import (
	"github.com/arafetki/go-echo-boilerplate/internal/app/api/chi/handler"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func (srv *server) routes(h *handler.Handler) {

	srv.router.Use(chiMiddleware.RequestID)
	srv.router.Use(chiMiddleware.RealIP)
	srv.router.Use(chiMiddleware.Logger)
	srv.router.Use(chiMiddleware.Recoverer)

	srv.router.Get("/health", h.HealthCheckHandler)
}
