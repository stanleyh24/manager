package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stanleyh24/manager/services"
)

func GetRouters() gin.HandlerFunc {
	return func(c *gin.Context) {
		routers, err := services.GetAllRouter()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food item"})
		}

		c.JSON(http.StatusOK, routers)
	}
}

func GetRouter() gin.HandlerFunc {
	return func(c *gin.Context) {

		/* routerId := c.Param("router_id")
		var router models.Router

		//err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)


		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food item"})
		}

		c.JSON(http.StatusOK, food) */
	}
}
