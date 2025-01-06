package middleware

import (
	"github.com/arafetki/go-echo-boilerplate/internal/config"
	"github.com/arafetki/go-echo-boilerplate/internal/logging"
)

type Middleware struct {
	cfg    config.Config
	logger logging.Logger
}

func New(cfg config.Config, logger logging.Logger) *Middleware {
	return &Middleware{
		cfg:    cfg,
		logger: logger,
	}
}
