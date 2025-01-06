package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/arafetki/go-echo-boilerplate/internal/utils"
)

func (app *Application) serveHTTP() error {

	app.echo.Server.ReadTimeout = app.config.Server.ReadTimeout
	app.echo.Server.WriteTimeout = app.config.Server.WriteTimeout

	shutdownErrChan := make(chan error)

	go func() {

		quitChan := make(chan os.Signal, 1)
		signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
		<-quitChan
		ctx, cancel := context.WithTimeout(context.Background(), app.config.Server.ShutdownPeriod)
		defer cancel()

		shutdownErrChan <- app.echo.Shutdown(ctx)

	}()

	app.Logger.Info("🚀 Server started", "env", utils.Capitalize(app.config.Env), "address", app.config.Server.Addr)
	if err := app.echo.Start(app.config.Server.Addr); err != nil && err != http.ErrServerClosed {
		return err
	}

	err := <-shutdownErrChan
	if err != nil {
		return err
	}
	app.wg.Wait()
	app.Logger.Warn("Server stopped gracefully")
	return nil
}
