package http

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/yael-castro/jikkosoft/3/internal/app/business"
	"net/http"
)

func ErrorHandler(handler echo.HTTPErrorHandler) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		var userErr business.Error

		if !errors.As(err, &userErr) {
			handler(err, c)
			return
		}

		code := http.StatusInternalServerError
		response := echo.Map{
			"code":    userErr.Error(),
			"message": err.Error(),
		}

		//goland:noinspection ALL
		switch userErr {
		case
			business.ErrEmptyOrder,
			business.ErrInvalidProduct,
			business.ErrInvalidStratum:
			code = http.StatusBadRequest
		}

		_ = c.JSON(code, response)
	}
}
