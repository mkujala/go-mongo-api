package controllers

import (
	"encoding/json"
	"fmt"
	"go-mongo-api/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductController struct{}

func NewProductController() *ProductController {
	return &ProductController{}
}

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

func (pc ProductController) Insert(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Waiting for implementation
}
