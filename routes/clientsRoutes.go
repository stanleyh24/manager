package routes

import (
	controllers "github.com/stanleyh24/manager/controllers"

	"github.com/gorilla/mux"
)

var RegisterClientRoutes = func(router *mux.Router) {
	router.HandleFunc("/client/", controllers.CreateClient).Methods("POST")
	router.HandleFunc("/client/", controllers.GetClients).Methods("GET")
	router.HandleFunc("/client/{clientId}", controllers.GetClientById).Methods("GET")
	router.HandleFunc("/client/{clientId}", controllers.UpdateClient).Methods("PUT")
	router.HandleFunc("/client/{clientId}", controllers.DeleteClient).Methods("DELETE")
}
