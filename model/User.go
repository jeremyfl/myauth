package model

import (
	"fmt"
	"myauth/config"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password" json:"password"`
	Level    string        `bson:"level" json:"level"`
}

var session, err = config.MongoConnect()

const (
	DOCUMENT = "user"
)

func GetUser() error {
	var user []User

	return nil
}

func InsertUser() error {
	var collection = session.DB(config.COLLECTION).C(DOCUMENT)

	err = collection.Insert(&User{"John", "jeremylombogia7@gmail.com", "rahasia", "admin"})
	if err != nil {
		fmt.Println("Error!", err.Error())
		return nil
	}

	return nil
}
