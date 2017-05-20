package main

import mgo "gopkg.in/mgo.v2"

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

}
