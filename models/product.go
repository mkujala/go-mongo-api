package models

import (
	"encoding/json"
	"go-mongo-api/config"
	"io/ioutil"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

// Product type includes product fields
type Product struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Image       string        `json:"image" bson:"image"`
	Price       float32       `json:"price" bson:"price"`
}

// AllProducts from MongoDB
func AllProducts() ([]Product, error) {
	prod := []Product{}
	err := config.Products.Find(bson.M{}).All(&prod)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

// PutProduct inserts to MongoDB
func PutProduct(r *http.Request) (Product, error) {
	prod := Product{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &prod)
	if err != nil {
		panic(err)
	}
	// fmt.Println(prod.Name)

	prod.ID = bson.NewObjectId()

	// insert values
	err = config.Products.Insert(prod)
	// if err != nil {
	// 	return prod, errors.New("500. Internal Server Error." + err.Error())
	// }
	return prod, err
}
