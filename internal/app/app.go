package app

import (
	"github.com/arafetki/go-echo-boilerplate/internal/app/api/rest"
	"github.com/arafetki/go-echo-boilerplate/internal/config"
	"github.com/arafetki/go-echo-boilerplate/internal/db"
	"github.com/arafetki/go-echo-boilerplate/internal/db/sqlc"
	"github.com/arafetki/go-echo-boilerplate/internal/logging"
	"github.com/arafetki/go-echo-boilerplate/internal/service"
)

type application struct {
	logger logging.Logger
	config config.Config
}

func New(cfg config.Config, logger logging.Logger) *application {
	return &application{
		logger: logger,
		config: cfg,
	}
}

func (app *application) Run() error {

	// Connect to database
	db, err := db.Pool(app.config.Database.Dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	app.logger.Info("Database connection established sucessfully")

	// Initialize services
	svc := service.New(sqlc.New(db))

	// Create server instance
	srv := rest.NewServer(app.config, app.logger, svc)

	// Start the server
	return srv.Start()
}
