package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

const SERVER = "localhost:27017"

func Get() *mgo.Session {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return session
}
