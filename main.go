package main

import (
	"go-mongo-api/product"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Constants
const (
	apiURL = "/api/v1"
	port   = ":8000"
)

func main() {
	router := httprouter.New()

	router.GET(apiURL+"/product", product.GetAll)
	router.POST(apiURL+"/product", product.Insert)
	router.DELETE(apiURL+"/product/:id", product.Delete)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("Server start failed when using PORT:", port)
	}
}
