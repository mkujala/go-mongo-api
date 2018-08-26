package main

import (
	"go-mongo-api/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const apiURL = "/api/v1"

func main() {
	router := httprouter.New()
	pc := controllers.NewProductController()

	router.GET(apiURL+"/product", pc.GetAll)
	router.POST(apiURL+"/product", pc.Insert)
	router.DELETE(apiURL+"/product/:id", pc.Delete)

	http.ListenAndServe(":8000", router)
}
