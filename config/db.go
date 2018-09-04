package config

import (
	"fmt"

	"github.com/globalsign/mgo"
)

// DB session
var DB *mgo.Database

// Products collection
var Products *mgo.Collection

func init() {
	const mongoURI = "mongodb://localhost"
	const mongoDB = "go_mongo_api"

	s, err := mgo.Dial(mongoURI)
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB(mongoDB)
	Products = DB.C("products")

	fmt.Println("You connected to your mongo database.")
}
