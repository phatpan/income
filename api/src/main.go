package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	Mongo_session    *mgo.Session
	Users_collection *mgo.Collection
)

const (
	API_SERVER                  = ":4200"
	DATABASE_SERVER             = "localhost:27017"
	DATABASE_NAME               = "account"
	DATABASE_COLLECTION_EXPENSE = "expense"
	DATABASE_COLLECTION_REVENUE = "revenue"
)

func main() {
	mongo_session, err := mgo.Dial(DATABASE_SERVER)
	if err != nil {
		panic(err)
	}

	mongo_session.SetMode(mgo.Monotonic, true)
	Mongo_session = mongo_session
	Users_collection = mongo_session.DB(DATABASE_NAME).C(DATABASE_COLLECTION_EXPENSE)
	defer Mongo_session.Close()

	err = Users_collection.Insert(bson.M{
		"amount": "1000",
	})
}
