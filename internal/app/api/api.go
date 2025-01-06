package api

import (
	"github.com/arafetki/go-echo-boilerplate/internal/app/api/rest"
	"github.com/arafetki/go-echo-boilerplate/internal/config"
	"github.com/arafetki/go-echo-boilerplate/internal/logging"
	"github.com/arafetki/go-echo-boilerplate/internal/service"
)

type api struct {
	Server interface {
		Start() error
	}
}

func New(cfg config.Config, logger logging.Logger, svc *service.Service) *api {
	return &api{
		Server: rest.NewServer(cfg, logger, svc),
	}
}
