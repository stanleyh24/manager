package routes

import (
	controller "github.com/stanleyh24/manager/api/controllers"

	"github.com/gin-gonic/gin"
)

func RouterRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/routers", controller.GetRouters())
	incomingRoutes.GET("/routers/:router_id", controller.GetRouter())
	/*incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id", controller.UpdateFood()) */
}
