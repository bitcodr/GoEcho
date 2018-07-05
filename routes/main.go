package routes

import (
	user "github.com/amiralii/goEchoExample/aggregates/user/route"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {

	user.Init(e)
}
