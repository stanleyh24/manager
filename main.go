package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stanleyh24/manager/database"
)

func main() {
	db := database.ConnectDB2()
	defer db.Close(context.Background())

	var id int
	var ip string
	err := db.QueryRow(context.Background(), "select id, ip from router where id=$1", 58).Scan(&id, &ip)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(id, ip)
}
