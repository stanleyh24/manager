package services

import (
	"fmt"

	"github.com/stanleyh24/manager/models"
	"github.com/stanleyh24/manager/repositories"
)

func GetAllRouter() ([]models.Router, error) {
	routers, err := repositories.Getall()

	if err != nil {
		fmt.Println("Error getting the router list ")
		return nil, err
	}

	return routers, nil
}

func GetRouterById(id int) (*models.Router, error) {
	router, err := repositories.Getbyid(id)

	if err != nil {
		fmt.Println("Error getting the router list ")
	}

	return &router, nil
}

func CreateRouter(r models.CreateRouterParams) (*models.Router, error) {

	return nil, nil
}

func UpdateRouter(id int, r models.Router) (*models.Router, error) {

	return nil, nil
}

func DeleteRouter(id int) error {

	return nil
}
