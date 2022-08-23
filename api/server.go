package api

/* import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome To the server")
}

func handleRequest() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Server() {
	handleRequest()
} */

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stanleyh24/manager/api/routes"
)

func Server() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	routes.RouterRoutes(router)
	router.Run(":" + port)
}
