package routes

import (
	controllers "github.com/stanleyh24/manager/controllers"

	"github.com/gorilla/mux"
)

var RegisterServiceRoutes = func(router *mux.Router) {
	router.HandleFunc("/service/", controllers.CreateService).Methods("POST")
	router.HandleFunc("/service/", controllers.GetServices).Methods("GET")
	router.HandleFunc("/service/{serviceId}", controllers.GetServiceById).Methods("GET")
	router.HandleFunc("/service/{serviceId}", controllers.UpdateService).Methods("PUT")
	router.HandleFunc("/service/{serviceId}", controllers.DeleteService).Methods("DELETE")
}
