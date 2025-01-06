package api

import (
	"net/http"

	"github.com/arafetki/go-echo-boilerplate/internal/app/api/handlers"
	"github.com/arafetki/go-echo-boilerplate/internal/app/api/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(r *echo.Echo, h *handlers.Handler, m *middlewares.Middleware) {

	r.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[ECHO] ${time_rfc3339} | ${method} | ${uri} | ${status} | ${latency_human} | ${remote_ip} | ${user_agent}\n",
	}))
	r.Use(middleware.Recover())
	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
	}))

	// health checks
	r.GET("/health", h.HealthCheckHandler)

	// API version 1 prefix
	v1 := r.Group("/v1")

	v1.Use(m.Authenticate)

	v1.POST("/users", h.CreateUserHandler)
	v1.GET("/users/:id", h.FetchUserDataHandler, m.RequireAuthenticatedUser)

}
