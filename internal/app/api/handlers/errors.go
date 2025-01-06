package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CustomHttpErrorHandler(err error, c echo.Context) {

	if c.Response().Committed {
		return
	}
	code := http.StatusInternalServerError
	var message any = "The server encountered a problem and could not process your request."
	if httpError, ok := err.(*echo.HTTPError); ok {
		code = httpError.Code
		switch code {
		case http.StatusNotFound:
			message = "The requested resource could not be found."
		case http.StatusMethodNotAllowed:
			message = fmt.Sprintf("The %s method is not supported for this resource.", c.Request().Method)
		case http.StatusBadRequest:
			message = "The request could not be understood or was missing required parameters."
		case http.StatusInternalServerError:
			message = "The server encountered a problem and could not process your request."
		case http.StatusUnprocessableEntity:
			message = "The request could not be processed due to invalid input."
		default:
			message = httpError.Message
		}
	} else {
		h.Logger.Error(err.Error())
	}
	if err := c.JSON(code, echo.Map{"message": message}); err != nil {
		h.Logger.Error(err.Error())
	}
}
