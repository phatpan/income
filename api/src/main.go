package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	mgo "gopkg.in/mgo.v2"
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

type Expense struct {
	Amount      string `json:"amount"`
	Description string `json:"description"`
}

func main() {
	mongo_session, err := mgo.Dial(DATABASE_SERVER)
	if err != nil {
		panic(err)
	}

	mongo_session.SetMode(mgo.Monotonic, true)
	Mongo_session = mongo_session
	Users_collection = mongo_session.DB(DATABASE_NAME).C(DATABASE_COLLECTION_EXPENSE)
	defer Mongo_session.Close()

	e := echo.New()
	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.POST("/expense", createExpense)
	e.Logger.Fatal(e.Start(API_SERVER))

	expense := Expense{
		Amount:      "500",
		Description: "ค่าเดินทาง",
	}
	expense.SaveToDB()
}

func createExpense(c echo.Context) error {

}

func (e *Expense) SaveToDB() error {
	err := Users_collection.Insert(&e)
	if err != nil {
		return err
	}
	return nil
}
