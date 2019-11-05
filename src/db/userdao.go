package db

import (
	"log"

	"github.com/akspokl/go-task/src/model"
	"gopkg.in/mgo.v2/bson"
)

const DBNAME = "userrepo"

const DOCNAME = "user"

func AddUser(user model.User) error {
	db := Get()
	defer db.Close()

	return db.DB(DBNAME).C(DOCNAME).Insert(user)
}

func GetAllUsers() ([]model.User, error) {
	db := Get()
	defer db.Close()

	res := []model.User{}
	log.Println(res)

	if err := db.DB(DBNAME).C(DOCNAME).Find(nil).All(&res); err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserByUsername(username string) (*model.User, error) {
	db := Get()
	defer db.Close()

	res := model.User{}

	if err := db.DB(DBNAME).C(DOCNAME).Find(bson.M{"username": username}).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}
