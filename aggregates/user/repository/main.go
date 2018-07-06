package userRepository

import (
	"os"

	"github.com/globalsign/mgo"

	"github.com/globalsign/mgo/bson"

	user "github.com/amiralii/goEchoExample/aggregates/user/model"
	"github.com/amiralii/goEchoExample/config/database"
	"github.com/amiralii/goEchoExample/config/response"
)

func NewUser(u user.User) error {
	db := database.DB{}
	session, err := db.Connect()
	if err != nil {
		return response.Error("Error on connecting to database", 1002)
	}
	defer session.Close()
	collection := session.DB(os.Getenv("MONGO_DATABASE")).C(user.UserCollection)

	err = collection.Insert(u)
	if err != nil {
		return response.Error("error to create new user", 1003)
	}
	return err
}

func FindUser(username string, password string) (user.User, error) {
	db := database.DB{}
	u := user.User{}
	session, err := db.Connect()
	if err != nil {
		return u, response.Error("Error on connecting to database", 1002)
	}
	defer session.Close()
	collection := session.DB(os.Getenv("MONGO_DATABASE")).C(user.UserCollection)

	err = collection.Find(bson.M{"username": username, "password": password}).One(&u)
	if err != nil {
		if err == mgo.ErrNotFound {
			return u, response.Error("invalid credtional", 1004)
		}
		return u, err
	}
	return u, err
}
