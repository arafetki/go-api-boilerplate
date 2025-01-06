package rest

import (
	"net/http"

	"github.com/arafetki/go-echo-boilerplate/internal/app/api/rest/handler"
	"github.com/arafetki/go-echo-boilerplate/internal/app/api/rest/middleware"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func (srv *server) routes(h *handler.Handler, m *middleware.Middleware) {

	srv.echo.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: "[ECHO] ${time_rfc3339} | ${method} | ${uri} | ${status} | ${latency_human} | ${remote_ip} | ${user_agent} | error: ${error}\n",
	}))
	srv.echo.Use(echoMiddleware.Recover())
	srv.echo.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
	}))

	// health checks
	srv.echo.GET("/health", h.HealthCheckHandler)

	// API version 1 prefix
	v1 := srv.echo.Group("/v1")

	v1.Use(m.Authenticate)

	v1.POST("/users", h.CreateUserHandler)
	v1.GET("/users/:id", h.FetchUserDataHandler, m.RequireAuthenticatedUser)

}
