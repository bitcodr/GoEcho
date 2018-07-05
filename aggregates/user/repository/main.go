package userRepository

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/globalsign/mgo"

	"github.com/globalsign/mgo/bson"

	user "github.com/amiralii/goEchoExample/aggregates/user/model"
	"github.com/amiralii/goEchoExample/config"
)

func NewUser(u user.User) error {
	db := config.DB{}
	session, err := db.DBDial()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Error on connecting to database"}
	}
	defer session.Close()
	collection := session.DB(db.DBName()).C(user.UserCollection)

	err = collection.Insert(u)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "error to create new user"}
	}
	return err
}

func FindUser(username string, password string) (user.User, error) {
	db := config.DB{}
	u := user.User{}
	session, err := db.DBDial()
	if err != nil {
		return u, &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Error on connecting to database"}
	}
	defer session.Close()
	collection := session.DB(db.DBName()).C(user.UserCollection)

	err = collection.Find(bson.M{"username": username, "password": password}).One(&u)
	if err != nil {
		if err == mgo.ErrNotFound {
			return u, &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid username or password"}
		}
		return u, err
	}
	return u, err
}
