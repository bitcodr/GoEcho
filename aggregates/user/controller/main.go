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
	if err := e.Bind(u); err != nil {
		return errors.New("error")
	}
	err := Repo.NewUser(*u)
	if err != nil {
		return errors.New("error on create user")
	}
	return e.JSON(http.StatusCreated, "user created")
}

func FindUser(e echo.Context) error {
	u := new(user.User)
	if err := e.Bind(u); err != nil {
		return errors.New("error")
	}
	if u.Username == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid username or password"}
	}
	response, err := Repo.FindUser(u.Username, u.Password)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, response)
}
