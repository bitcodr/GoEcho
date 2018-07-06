package response

import (
	"net/http"

	"github.com/labstack/echo"
)

type message struct {
	message   string
	errorCode int
}

// func ok() {

// }

func Error(text string, code int) error {
	return &echo.HTTPError{Code: http.StatusUnprocessableEntity, Message: &message{message: text, errorCode: code}}
}
