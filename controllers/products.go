package controllers

import (
	"encoding/json"
	"fmt"
	"go-mongo-api/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ProductController for handling product data
type ProductController struct{}

// NewProductController returns pointer to ProductController
func NewProductController() *ProductController {
	return &ProductController{}
}

// GetAll products from DB
func (pc ProductController) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	prods, err := models.AllProducts()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal into JSON
	pj, err := json.Marshal(prods)
	if err != nil {
		fmt.Println(err)
	}

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", pj)
}

// Insert new products to DB
func (pc ProductController) Insert(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Waiting for implementation

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	if r.Body == nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// ADD HANDLING FOR ARRAY of JSON

	prod, err := models.PutProduct(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	prodj, err := json.Marshal(prod)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", prodj)
}

// Delete product from DB
func (pc ProductController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	if p.ByName("id") == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	err := models.DeleteProduct(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent) // 204
}
