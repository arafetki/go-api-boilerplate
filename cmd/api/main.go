package main

import (
	"log/slog"
	"os"
	"runtime/debug"

	"github.com/arafetki/go-echo-boilerplate/internal/app"
	"github.com/arafetki/go-echo-boilerplate/internal/config"
	"github.com/arafetki/go-echo-boilerplate/internal/logging"
)

func main() {

	// Initialize the configuration
	cfg := config.Init()

	// Set the log level based on the debug flag
	slogLevel := slog.LevelInfo
	if cfg.Debug {
		slogLevel = slog.LevelDebug
	}

	// Create slog instance
	slogLogger := logging.NewSlogLogger(os.Stdout, slogLevel)

	// Initialize the application
	app := app.New(cfg, slogLogger)

	// Run the application
	if err := app.Run(); err != nil {
		trace := string(debug.Stack())
		slogLogger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}
