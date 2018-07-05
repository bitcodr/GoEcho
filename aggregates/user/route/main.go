package userRoute

import (
	user "github.com/amiralii/goEchoExample/aggregates/user/controller"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	e.POST("/signup", user.NewUser)
	e.POST("/login", user.FindUser)
}
