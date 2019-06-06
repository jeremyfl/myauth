package model

import (
	"fmt"
	"myauth/config"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	Level    string `bson:"level" json:"level"`
}

var session, _ = config.MongoConnect()

const DOCUMENT string = "user"

var collection = session.DB(config.COLLECTION).C(DOCUMENT)

func InsertUser(user *User) error {
	var err = collection.Insert(&user)
	if err != nil {
		fmt.Println("Error", err.Error())
		return nil
	}

	return err
}

func GetUser() ([]User, error) {
	var user []User

	var err = collection.Find(bson.M{}).All(&user)

	return user, err
}
