package main

import (
	"os"

	"github.com/amiralii/goEchoExample/config/env"
	"github.com/amiralii/goEchoExample/routes"
	"github.com/labstack/echo"
)

func main() {

	env.Init()

	e := echo.New()

	routes.Init(e)

	e.Start(os.Getenv("APP_PORT"))
}
