package response

import (
	"net/http"

	"github.com/labstack/echo"
)

type errorMessage struct {
	message   string
	errorCode int
}

type created struct {
	data string
}

func Created(e echo.Context, data string) error {
	return e.JSON(http.StatusCreated, &created{data: data})
}

func Ok(e echo.Context, data interface{}) error {
	return e.JSON(http.StatusOK, data)
}

func Error(text string, code int) error {
	return &echo.HTTPError{Code: http.StatusUnprocessableEntity, Message: &errorMessage{message: text, errorCode: code}}
}
