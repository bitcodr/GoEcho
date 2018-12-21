package response

import (
	"net/http"

	"github.com/labstack/echo"
)

type errorMessage struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"code"`
}

type errorResponse struct {
	error errorMessage
}

type created struct {
	data map[string]string
}

func Created(e echo.Context, message string) error {
	response := &created{map[string]string{"data": message}}
	return e.JSON(http.StatusCreated, response.data)
}

func Ok(e echo.Context, data interface{}) error {
	return e.JSON(http.StatusOK, data)
}

func Error(text string, code int) error {
	response := &errorResponse{errorMessage{Message: text, ErrorCode: code}}
	return &echo.HTTPError{Code: http.StatusUnprocessableEntity, Message: response.error}
}
