package userController

import (
	"net/http"

	user "github.com/amiralii/goEchoExample/aggregates/user/model"
	Repo "github.com/amiralii/goEchoExample/aggregates/user/repository"
	"github.com/labstack/echo"
)

func Signup(e echo.Context) error {
	u := new(user.User)
	e.Bind(u)
	err := Repo.NewUser(*u)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusCreated, "user created")
}

func Signin(e echo.Context) error {
	u := new(user.User)
	e.Bind(u)
	if u.Username == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid username or password"}
	}
	response, err := Repo.FindUser(u.Username, u.Password)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, response)
}
