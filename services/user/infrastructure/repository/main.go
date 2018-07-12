package userRepository

import (
	"github.com/globalsign/mgo"

	"github.com/globalsign/mgo/bson"

	"github.com/amiralii/goEchoExample/config/database"
	"github.com/amiralii/goEchoExample/config/response"
	user "github.com/amiralii/goEchoExample/services/user/domain/model"
)

func Save(u user.UserSingup) error {
	collection, _ := database.Init(user.UserCollection)
	err := collection.Insert(u)
	if err != nil {
		return response.Error("error to create new user", 1003)
	}
	return err
}

func Get(username string, password string) (user.UserSignin, error) {
	collection, _ := database.Init(user.UserCollection)
	u := user.UserSignin{}
	err := collection.Find(bson.M{"username": username, "password": password}).One(&u)
	if err != nil {
		if err == mgo.ErrNotFound {
			return u, response.Error("invalid credtional", 1004)
		}
		return u, err
	}
	return u, err
}
