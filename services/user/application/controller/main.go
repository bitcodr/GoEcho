package userController

import (
	"github.com/amiralii/goEchoExample/config/response"

	user "github.com/amiralii/goEchoExample/services/user/domain/model"
	Repo "github.com/amiralii/goEchoExample/services/user/infrastructure/repository"
	"github.com/labstack/echo"
)

func Signup(e echo.Context) error {
	u := new(user.UserSingup)
	if err := e.Bind(u); err != nil {
		return response.Error(err.Error(), 500)
	}
	if err := e.Validate(u); err != nil {
		return response.Error(err.Error(), 422)
	}
	if err := Repo.Save(*u); err != nil {
		return response.Error(err.Error(), 500)
	}
	return response.Created(e, "user created")
}

func Signin(e echo.Context) error {
	u := new(user.UserSignin)
	if err := e.Bind(u); err != nil {
		return response.Error(err.Error(), 500)
	}
	data, err := Repo.Get(u.Username, u.Password)
	if err != nil {
		return response.Error(err.Error(), 500)
	}
	return response.Ok(e, data)
}
