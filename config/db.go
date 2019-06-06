package config

import (
	"gopkg.in/mgo.v2"
)

const (
	CONNECTION_NAME = "localhost"
	COLLECTION      = "myauth"
)

func MongoConnect() (*mgo.Session, error) {
	var session, err = mgo.Dial(CONNECTION_NAME)
	if err != nil {
		return nil, err
	}

	return session, nil
}
