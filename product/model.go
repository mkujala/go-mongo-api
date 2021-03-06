package product

import (
	"encoding/json"
	"fmt"
	"go-mongo-api/config"
	"io/ioutil"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

// Product type includes product fields
type Product struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name,omitempty" bson:"name,omitempty"` // required
	Description string        `json:"description" bson:"description"`
	Image       string        `json:"image" bson:"image"`
	Price       float32       `json:"price" bson:"price"`
}

// allProducts from MongoDB
func allProducts() ([]Product, error) {
	prod := []Product{}
	err := config.Products.Find(bson.M{}).All(&prod)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

// filterProducts with search string
func filterProducts(field string, search string) ([]Product, error) {
	//-----------------
	// WORK IN PROGRESS
	//-----------------
	prod := []Product{}
	err := config.Products.Find(bson.M{}).All(&prod)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

// putProduct inserts to MongoDB
func putProduct(r *http.Request) (Product, error) {
	prod := Product{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return prod, err
	}
	// fmt.Printf("type %T, value %v\n", body, string(body)) // DEBUG print request body

	err = json.Unmarshal(body, &prod) // -> prod.Name, prod.Image, prod.Price etc.
	if err != nil {
		return prod, err
	}
	prod.ID = bson.NewObjectId()

	// Check for required fields with Validator
	var v Validator
	if !v.required(prod.Name) {
		err = fmt.Errorf("Field 'name' is required")
		return prod, err
	}

	// insert values
	err = config.Products.Insert(prod)
	return prod, err
}

// deleteProduct by _id from MongoDB
func deleteProduct(id string) error {
	bsonObjectID := bson.ObjectIdHex(id)
	err := config.Products.Remove(bson.M{"_id": bsonObjectID})
	if err != nil {
		return err
	}
	return nil
}
