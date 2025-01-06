package handlers

import (
	"github.com/arafetki/go-echo-boilerplate/internal/logging"
	"github.com/arafetki/go-echo-boilerplate/internal/service"
)

type Handler struct {
	Service *service.Service
	Logger  logging.Logger
}
