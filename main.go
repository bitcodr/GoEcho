package main

import (
	"os"

	"github.com/amiralii/goEchoExample/config/validation"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/amiralii/goEchoExample/config/env"
	"github.com/amiralii/goEchoExample/routing"
	"github.com/labstack/echo"
)

func main() {

	env.Init()

	e := echo.New()

	e.Validator = &validation.DataValidator{ValidatorData: validator.New()}

	routing.Init(e)

	e.Start(os.Getenv("APP_PORT"))
}
