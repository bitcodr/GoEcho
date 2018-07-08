package database

import (
	"os"

	"github.com/amiralii/goEchoExample/config/response"
	"github.com/globalsign/mgo"
)

type dbSession struct {
	Session *mgo.Session
}

func (db *dbSession) dial() (s *mgo.Session, err error) {
	return mgo.Dial(os.Getenv("MONGO_HOST"))
}

func Init(collection string) (s *mgo.Collection, err error) {
	db := &dbSession{}
	session, err := db.dial()
	if err != nil {
		return nil, response.Error("database connection error", 100)
	}
	// defer session.Close()
	return session.DB(os.Getenv("MONGO_DATABASE")).C(collection), nil
}
