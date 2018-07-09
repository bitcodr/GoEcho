package response

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type errorMessage struct {
	message   map[string]string
	errorCode map[string]int
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
	response := &errorMessage{map[string]string{"message": text}, map[string]int{"errorCode": code}}
	fmt.Println(response)
	return &echo.HTTPError{Code: http.StatusUnprocessableEntity, Message: response}
}
