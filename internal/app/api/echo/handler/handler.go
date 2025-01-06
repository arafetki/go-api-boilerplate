package handler

import (
	"github.com/arafetki/go-api-boilerplate/internal/logging"
	"github.com/arafetki/go-api-boilerplate/internal/service"
)

type Handler struct {
	service *service.Service
	logger  logging.Logger
}

func New(service *service.Service, logger logging.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}
