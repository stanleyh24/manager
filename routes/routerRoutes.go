package routes

import (
	controllers "github.com/stanleyh24/manager/controllers"

	"github.com/gorilla/mux"
)

var RegisterRouterRoutes = func(router *mux.Router) {
	router.HandleFunc("/router/", controllers.CreateRouter).Methods("POST")
	router.HandleFunc("/router/", controllers.GetRouters).Methods("GET")
	router.HandleFunc("/router/{routerId}", controllers.GetRouterById).Methods("GET")
	router.HandleFunc("/router/{routerId}", controllers.UpdateRouter).Methods("PUT")
	router.HandleFunc("/router/{routerId}", controllers.DeleteRouter).Methods("DELETE")
}
