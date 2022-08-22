package services

import (
	"fmt"

	"github.com/stanleyh24/manager/models"
	"github.com/stanleyh24/manager/repositories"
)

func GetAllRouter() (models.Routers, error) {
	routers, err := repositories.Getall()

	if err != nil {
		fmt.Println("Error getting the router list: ")
		return nil, err
	}

	return routers, nil
}
