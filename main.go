package main

import (
	"fmt"
	"log"

	"github.com/stanleyh24/manager/api"
	"github.com/stanleyh24/manager/services"
)

func main() {
	/* db := database.ConnectDB()
	defer db.Close(context.Background())

	var id int
	var ip string
	err := db.QueryRow(context.Background(), "select id, ip from router where id=$1", 58).Scan(&id, &ip)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(id, ip) */

	routers, err := services.GetAllRouter()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(routers)

	api.Server()
}
