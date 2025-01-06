package middlewares

import (
	"github.com/arafetki/go-echo-boilerplate/internal/config"
	"github.com/arafetki/go-echo-boilerplate/internal/logging"
)

type Middleware struct {
	Config config.Config
	Logger logging.Logger
}
