package database

import (
	"os"

	"github.com/globalsign/mgo"
)

type DB struct {
	Session *mgo.Session
}

func (db *DB) Connect() (s *mgo.Session, err error) {
	return mgo.Dial(os.Getenv("MONGO_HOST"))
}
