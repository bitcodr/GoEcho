package userController

import (
	"errors"
	"net/http"

	user "github.com/amiralii/goEchoExample/aggregates/user/model"
	Repo "github.com/amiralii/goEchoExample/aggregates/user/repository"
	"github.com/labstack/echo"
)

func NewUser(e echo.Context) error {
	u := new(user.User)
	e.Bind(u)
	err := Repo.NewUser(*u)
	if err != nil {
		return errors.New("error on create user")
	}
	return e.JSON(http.StatusCreated, "user created")
}

func FindUser(e echo.Context) error {
	u := new(user.User)
	e.Bind(u)
	response, err := Repo.FindUser(u.Username, u.Password)
	if err != nil {
		return errors.New("error to find user")
	}
	return e.JSON(http.StatusOK, response)
}
