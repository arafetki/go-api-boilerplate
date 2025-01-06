package app

import (
	"log/slog"
	"os"
	"sync"

	"github.com/arafetki/go-echo-boilerplate/internal/app/api"
	"github.com/arafetki/go-echo-boilerplate/internal/app/api/handlers"
	"github.com/arafetki/go-echo-boilerplate/internal/app/api/middlewares"
	"github.com/arafetki/go-echo-boilerplate/internal/config"
	"github.com/arafetki/go-echo-boilerplate/internal/db"
	"github.com/arafetki/go-echo-boilerplate/internal/db/sqlc"
	"github.com/arafetki/go-echo-boilerplate/internal/logging"
	"github.com/arafetki/go-echo-boilerplate/internal/service"
	"github.com/arafetki/go-echo-boilerplate/internal/validator"
	"github.com/labstack/echo/v4"
)

type Application struct {
	echo   *echo.Echo
	config config.Config
	Logger logging.Logger
	wg     sync.WaitGroup
}

func Init() *Application {

	// Initialize the configuration
	cfg := config.Init()

	// Set the log level based on the debug flag
	slogLevel := slog.LevelInfo
	if cfg.Debug {
		slogLevel = slog.LevelDebug
	}
	slogLogger := logging.NewSlogLogger(os.Stdout, slogLevel)

	// Initialize the Echo instance
	e := echo.New()
	e.Debug = cfg.Debug
	e.HideBanner = !cfg.Debug
	e.HidePort = true
	e.Validator = validator.New()

	// Return the application struct with initialized components
	return &Application{
		echo:   e,
		config: cfg,
		Logger: slogLogger,
	}
}

func (app *Application) Run() error {

	// Connect to database
	db, err := db.Pool(app.config.Database.Dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	app.Logger.Info("Database connection established sucessfully")

	// Initialize services
	svc := service.New(sqlc.New(db))

	// Initialize HTTP handler.
	handler := &handlers.Handler{
		Service: svc,
		Logger:  app.Logger,
	}

	// Initialize Middleware.
	middleware := &middlewares.Middleware{
		Config: app.config,
	}

	// Register routes
	api.RegisterRoutes(app.echo, handler, middleware)

	return app.serveHTTP()

}
