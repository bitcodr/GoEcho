package userController

import (
	"net/http"

	"github.com/amiralii/goEchoExample/config/response"

	user "github.com/amiralii/goEchoExample/aggregates/user/model"
	Repo "github.com/amiralii/goEchoExample/aggregates/user/repository"
	"github.com/labstack/echo"
)

func Signup(e echo.Context) error {
	u := new(user.User)
	if err := e.Bind(u); err != nil {
		return err
	}
	if err := e.Validate(u); err != nil {
		return response.Error(err.Error(), 422)
	}
	err := Repo.Signup(*u)
	if err != nil {
		return err
	}
	return response.Created(e, "user created")
}

func Signin(e echo.Context) error {
	u := new(user.User)
	e.Bind(u)
	if u.Username == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid username or password"}
	}
	data, err := Repo.FindUser(u.Username, u.Password)
	if err != nil {
		return err
	}
	return response.Ok(e, data)
}
