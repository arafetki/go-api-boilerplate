package middlewares

import (
	"net/http"

	"github.com/arafetki/go-echo-boilerplate/internal/utils"
	"github.com/labstack/echo/v4"
)

func (m *Middleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Add("Vary", "Authorization")

		r := utils.ContextSetUser(c.Request(), &utils.DummyUser{
			ID: 2,
		})

		c.SetRequest(r)
		return next(c)
	}
}

func (m *Middleware) RequireAuthenticatedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := utils.ContextGetUser(c.Request())
		if user.IsAnonymous() {
			return echo.NewHTTPError(http.StatusUnauthorized, "You must be authenticated to access this resource")
		}

		return next(c)
	}
}
