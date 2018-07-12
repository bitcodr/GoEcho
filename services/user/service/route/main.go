package userRoute

import (
	user "github.com/amiralii/goEchoExample/services/user/application/controller"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {

	route := e.Group("/api/v1_0/users")
	{
		route.POST("/signup", user.Signup)
		route.POST("/signin", user.Signin)
		// e.DELETE("/signout", user.Signout)
		// e.POST("/forgot", user.ForgotPassword)
		// e.POST("/resert", user.ResetPassword)
		// e.GET("/sessions", user.Sessions)
	}

}
