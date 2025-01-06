package handler

import (
	"errors"
	"net/http"

	"github.com/arafetki/go-api-boilerplate/internal/db/sqlc"
	"github.com/arafetki/go-api-boilerplate/internal/service"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateUserHandler(c echo.Context) error {
	var input struct {
		Email string `json:"email" validate:"required,email"`
		Name  string `json:"name" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	err := h.service.Users.Create(sqlc.CreateUserParams{
		Email: input.Email,
		Name:  input.Name,
	})

	if err != nil {
		if errors.Is(err, service.ErrDuplicateEmail) {
			return echo.NewHTTPError(http.StatusConflict, "The email address provided is already associated with an existing account.")
		}
		h.logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)

}

func (h *Handler) FetchUserDataHandler(c echo.Context) error {
	var params struct {
		ID int32 `param:"id"`
	}
	if err := c.Bind(&params); err != nil {
		h.logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	user, err := h.service.Users.GetOne(params.ID)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		h.logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": user})
}
