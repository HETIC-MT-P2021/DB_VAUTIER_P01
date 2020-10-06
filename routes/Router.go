package routes

import (
	"github.com/edwinvautier/DB_VAUTIER_P01/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

	api := router.Group("/api")
	{
		api.GET("/", controllers.SayHello)

		v1 := api.Group("/v1")
		{
			v1.GET("/customers/:number", controllers.GetCustomerByNumber)
			v1.GET("/customers/:number/orders", controllers.GetOrdersByCustomerNumber)
			v1.GET("/orders/:number", controllers.GetOrderByNumber)
			v1.GET("/orders/:number/details", controllers.GetOrderDetailsByOrderNumber)
			v1.GET("/offices/:code/employees", controllers.GetEmployeesByOfficeCode)
		}
	}
}
