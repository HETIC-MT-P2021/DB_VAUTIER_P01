package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/edwinvautier/DB_VAUTIER_P01/controllers"
)

func SetupRouter(router *gin.Engine) {
		
	api := router.Group("/api")
	{
		api.GET("/", controllers.SayHello)

		v1 := api.Group("/v1")
		{
			v1.GET("/customers/:id", controllers.GetCustomerByNumber)
		}
	}
}