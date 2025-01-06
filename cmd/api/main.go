package main

import (
	"log/slog"
	"os"
	"runtime/debug"

	"github.com/arafetki/go-echo-boilerplate/internal/app/api"
	"github.com/arafetki/go-echo-boilerplate/internal/app/api/echo"
	"github.com/arafetki/go-echo-boilerplate/internal/config"
	"github.com/arafetki/go-echo-boilerplate/internal/db"
	"github.com/arafetki/go-echo-boilerplate/internal/db/sqlc"
	"github.com/arafetki/go-echo-boilerplate/internal/logging"
	"github.com/arafetki/go-echo-boilerplate/internal/service"
)

func main() {

	// Create slog instance
	logger := logging.NewSlogLogger(os.Stdout, slog.LevelInfo)

	// Run the application
	if err := run(logger); err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

func run(logger logging.SlogLogger) error {
	// Initialize the configuration
	cfg := config.Init()

	// Set the log level based on the debug flag
	if cfg.Debug {
		logger.SetLevel(slog.LevelDebug)
	}

	// Connect to database
	db, err := db.Pool(cfg.Database.Dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	logger.Info("Database connection established sucessfully")

	// Initialize services
	svc := service.New(sqlc.New(db))

	// Initialize Echo server instance
	server := echo.NewServer(cfg, logger, svc)

	// Initialize API instance
	api := api.New(server)

	return api.Server.Start()
}
