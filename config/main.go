package config

import (
	"os"

	"github.com/globalsign/mgo"
)

type DB struct {
	Session *mgo.Session
}

func (db *DB) DBDial() (s *mgo.Session, err error) {
	return mgo.Dial(dbURL())
}

func (db *DB) DBName() string {
	return "echoTest"
}

func dbURL() string {
	url := os.Getenv("MONGOHQ_URL")
	if url == "" {
		return "127.0.0.1:27017"
	}
	return url
}

func ServerPort() string {
	return ":8080"
}
