package dao

import (
	"golang-web-project-framework/src/utils/db/mongo"
)

const (
	databaseName   = "user"
	collectionName = "user"
)

var UserDaoMgo *UserDaoMongo

type UserDaoMongo struct {
	Mgo *mongo.Mgo
}

func init() {
	UserDaoMgo = &UserDaoMongo{
		Mgo: mongo.NewMgo(databaseName, collectionName),
	}
}
