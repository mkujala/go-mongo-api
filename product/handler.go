package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetAll products from DB
func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	prods, err := allProducts()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}

	pjson, err := json.Marshal(prods)
	if err != nil {
		fmt.Println(err)
	}

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", pjson)
}

// Insert new products to DB
func Insert(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Body == nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// ADD HANDLING FOR ARRAY of JSON
	//-----------------
	// WORK IN PROGRESS
	//-----------------

	prod, err := putProduct(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pjson, err := json.Marshal(prod)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", pjson)
}

// Delete product from DB
func Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if p.ByName("id") == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	err := deleteProduct(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent) // 204
}
