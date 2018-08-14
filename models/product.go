package models

import (
	"go-mongo-api/config"

	"github.com/globalsign/mgo/bson"
)

type Product struct {
	ID          bson.ObjectId // `json:"id" bson:"_id"`
	Name        string        // `json:"name" bson:"name"`
	Description string        // `json:"description" bson:"description"`
	Image       string        // `json:"image" bson:"image"`
	Price       float32       // `json:"price" bson:"price"`
}

func AllProducts() ([]Product, error) {
	prod := []Product{}
	err := config.Products.Find(bson.M{}).All(&prod)
	if err != nil {
		return nil, err
	}
	return prod, nil
}
