package routing

import (
	user "github.com/amiralii/goEchoExample/services/user/service/route"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {

	user.Init(e)
}
