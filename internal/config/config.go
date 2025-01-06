package config

import (
	"fmt"
	"time"

	"github.com/arafetki/go-echo-boilerplate/internal/env"
)

type Config struct {
	Env    string
	Debug  bool
	Server struct {
		Addr           string
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
		ShutdownPeriod time.Duration
	}
	Database struct {
		Dsn string
	}
}

func Init() Config {
	var cfg Config

	cfg.Env = env.GetString("APP_ENV", "development")
	cfg.Debug = env.GetBool("APP_DEBUG", true)
	cfg.Server.Addr = fmt.Sprintf(":%d", env.GetInt("SERVER_PORT", 8080))
	cfg.Server.ReadTimeout = env.GetDuration("SERVER_READ_TIMEOUT", 10*time.Second)
	cfg.Server.WriteTimeout = env.GetDuration("SERVER_WRITE_TIMEOUT", 30*time.Second)
	cfg.Server.ShutdownPeriod = env.GetDuration("SERVER_SHUTDOWN_PERIOD", 60*time.Second)
	cfg.Database.Dsn = env.GetString("DATABASE_DSN", "postgres:postgres@localhost:5432/test?sslmode=disable")
	return cfg
}
