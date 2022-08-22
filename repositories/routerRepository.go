package repositories

import (
	"context"
	"fmt"
	"os"

	"github.com/stanleyh24/manager/models"

	"github.com/stanleyh24/manager/database"
)

/* type router interface {
	Getall() models.Routers
	Getbyid(id int) (*models.Router, error)
	Create(r models.Createrouter) (*models.Router, error)
	Update(id int, r models.Router) (*models.Router, error)
	Delete(id int) error
} */

func Getall() (models.Routers, error) {

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

	return nil, nil
}

func Getbyid(id int) (*models.Router, error) {
	return nil, nil
}

func Create(r models.Createrouter) (*models.Router, error) {
	return nil, nil
}

func Update(id int, r models.Router) (*models.Router, error) {
	return nil, nil
}

func Delete(id int) error {
	return nil
}
