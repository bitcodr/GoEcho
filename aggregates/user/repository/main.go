package userRepository

import (
	"errors"

	"github.com/globalsign/mgo/bson"

	user "github.com/amiralii/goEchoExample/aggregates/user/model"
	"github.com/amiralii/goEchoExample/config"
)

func NewUser(u user.User) error {
	db := config.DB{}
	session, err := db.DBDial()
	if err != nil {
		return errors.New("database connection error")
	}
	defer session.Close()
	collection := session.DB(db.DBName()).C(user.UserCollection)

	err = collection.Insert(u)

	if err != nil {
		return errors.New("error on create new user")
	}
	return err
}

func FindUser(username string, password string) (user.User, error) {
	db := config.DB{}
	u := user.User{}
	session, err := db.DBDial()
	if err != nil {
		return u, errors.New("database connection error")
	}
	defer session.Close()
	collection := session.DB(db.DBName()).C(user.UserCollection)

	err = collection.Find(bson.M{"username": username, "password": password}).One(&u)
	if err != nil {
		return u, errors.New("error on find user")
	}
	return u, err
}
