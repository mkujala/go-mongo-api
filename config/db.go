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
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("go_mongo_api")
	Products = DB.C("products")

	fmt.Println("You connected to your mongo database.")
}
