package main

import (
	"github.com/amiralii/goEchoExample/config"
	"github.com/amiralii/goEchoExample/routes"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	routes.Init(e)

	e.Start(config.ServerPort())
}
