package rest

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/arafetki/go-echo-boilerplate/internal/app/api/rest/handler"
	"github.com/arafetki/go-echo-boilerplate/internal/app/api/rest/middleware"
	"github.com/arafetki/go-echo-boilerplate/internal/app/api/rest/validator"
	"github.com/arafetki/go-echo-boilerplate/internal/config"
	"github.com/arafetki/go-echo-boilerplate/internal/logging"
	"github.com/arafetki/go-echo-boilerplate/internal/service"
	"github.com/arafetki/go-echo-boilerplate/internal/utils"
	"github.com/labstack/echo/v4"
)

type server struct {
	echo    *echo.Echo
	logger  logging.Logger
	cfg     config.Config
	service *service.Service
	wg      sync.WaitGroup
}

func NewServer(cfg config.Config, logger logging.Logger, svc *service.Service) *server {
	server := &server{
		echo:    echo.New(),
		logger:  logger,
		cfg:     cfg,
		service: svc,
	}

	// Configure the server
	server.configure()

	// Register routes
	server.routes(
		handler.New(server.service, server.logger),
		middleware.New(server.cfg, server.logger),
	)

	return server
}

func (srv *server) configure() {
	srv.echo.Debug = srv.cfg.Debug
	srv.echo.HideBanner = !srv.cfg.Debug
	srv.echo.HidePort = true
	srv.echo.Validator = validator.New()
	srv.echo.Server.ReadTimeout = srv.cfg.Server.ReadTimeout
	srv.echo.Server.WriteTimeout = srv.cfg.Server.WriteTimeout
	srv.echo.HTTPErrorHandler = handleErrors(srv.logger)
}

func (srv *server) Start() error {

	shutdownErrChan := make(chan error)

	go func() {

		quitChan := make(chan os.Signal, 1)
		signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
		<-quitChan
		ctx, cancel := context.WithTimeout(context.Background(), srv.cfg.Server.ShutdownPeriod)
		defer cancel()

		shutdownErrChan <- srv.echo.Shutdown(ctx)

	}()

	srv.logger.Info("🚀 Server started", "env", utils.Capitalize(srv.cfg.Env), "address", srv.cfg.Server.Addr)
	if err := srv.echo.Start(srv.cfg.Server.Addr); err != nil && err != http.ErrServerClosed {
		return err
	}

	err := <-shutdownErrChan
	if err != nil {
		return err
	}
	srv.wg.Wait()
	srv.logger.Warn("Server stopped gracefully")
	return nil
}

func handleErrors(logger logging.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}
		code := http.StatusInternalServerError
		var message any = "The server encountered a problem and could not process your request."
		if httpError, ok := err.(*echo.HTTPError); ok {
			code = httpError.Code
			switch code {
			case http.StatusNotFound:
				message = "The requested resource could not be found."
			case http.StatusMethodNotAllowed:
				message = fmt.Sprintf("The %s method is not supported for this resource.", c.Request().Method)
			case http.StatusBadRequest:
				message = "The request could not be understood or was missing required parameters."
			case http.StatusInternalServerError:
				message = "The server encountered a problem and could not process your request."
			case http.StatusUnprocessableEntity:
				message = "The request could not be processed due to invalid input."
			default:
				message = httpError.Message
			}
		} else {
			logger.Error(err.Error())
		}
		if err := c.JSON(code, echo.Map{"message": message}); err != nil {
			logger.Error(err.Error())
		}
	}
}
