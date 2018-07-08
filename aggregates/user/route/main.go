package userRoute

import (
	user "github.com/amiralii/goEchoExample/aggregates/user/controller"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {

	e.Group("/api/v1_0/users")
	{
		e.POST("/signup", user.Signup)
		e.POST("/signin", user.Signin)
		// e.DELETE("/signout", user.Signout)
		// e.POST("/forgot", user.ForgotPassword)
		// e.POST("/resert", user.ResetPassword)
		// e.GET("/sessions", user.Sessions)
	}

}
