package handlers

import (
	"errors"
	"net/http"

	"github.com/arafetki/go-echo-boilerplate/internal/db/sqlc"
	"github.com/arafetki/go-echo-boilerplate/internal/service"
	"github.com/arafetki/go-echo-boilerplate/internal/utils"
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

	err := h.Service.Users.Create(sqlc.CreateUserParams{
		Email: input.Email,
		Name:  input.Name,
	})

	if err != nil {
		if errors.Is(err, service.ErrDuplicateEmail) {
			return echo.NewHTTPError(http.StatusConflict, "The email address provided is already associated with an existing account.")
		}
		h.Logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)

}

func (h *Handler) FetchUserDataHandler(c echo.Context) error {
	var params struct {
		ID int32 `param:"id"`
	}
	if err := c.Bind(&params); err != nil {
		h.Logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	authenticatedUser := utils.ContextGetUser(c.Request())
	if authenticatedUser.ID != params.ID {
		return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to perform this action.")
	}

	user, err := h.Service.Users.GetOne(params.ID)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		h.Logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": user})
}
