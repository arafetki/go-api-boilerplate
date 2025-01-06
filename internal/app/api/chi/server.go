package chi

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/arafetki/go-api-boilerplate/internal/app/api/chi/handler"
	"github.com/arafetki/go-api-boilerplate/internal/config"
	"github.com/arafetki/go-api-boilerplate/internal/logging"
	"github.com/arafetki/go-api-boilerplate/internal/service"
	"github.com/go-chi/chi/v5"
)

type server struct {
	router  *chi.Mux
	logger  logging.Logger
	cfg     config.Config
	service *service.Service
	wg      sync.WaitGroup
}

func NewServer(cfg config.Config, logger logging.Logger, svc *service.Service) *server {
	server := &server{
		router:  chi.NewRouter(),
		logger:  logger,
		cfg:     cfg,
		service: svc,
	}

	server.routes(handler.New(svc, logger))

	return server
}

func (srv *server) Start() error {

	server := &http.Server{
		Addr:         srv.cfg.Server.Addr,
		Handler:      srv.router,
		ReadTimeout:  srv.cfg.Server.ReadTimeout,
		WriteTimeout: srv.cfg.Server.WriteTimeout,
	}

	shutdownErrChan := make(chan error)

	go func() {

		quitChan := make(chan os.Signal, 1)
		signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
		<-quitChan
		ctx, cancel := context.WithTimeout(context.Background(), srv.cfg.Server.ShutdownPeriod)
		defer cancel()

		shutdownErrChan <- server.Shutdown(ctx)

	}()

	srv.logger.Info("🚀 server started", "env", srv.cfg.Env, "address", srv.cfg.Server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	err := <-shutdownErrChan
	if err != nil {
		return err
	}
	srv.wg.Wait()
	srv.logger.Warn("server stopped gracefully")
	return nil
}
