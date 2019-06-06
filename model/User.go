package model

import (
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

var session, _ = config.MongoConnect()

const DOCUMENT string = "user"

var collection = session.DB(config.COLLECTION).C(DOCUMENT)

func InsertUser(user *User) error {
	return collection.Insert(&user)
}

func GetUser() ([]User, error) {
	var user []User

	var err = collection.Find(bson.M{}).All(&user)

	return user, err
}
