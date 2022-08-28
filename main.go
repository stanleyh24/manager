package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stanleyh24/manager/routes"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	routes.RegisterRouterRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))

}
