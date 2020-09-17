package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/edwinvautier/bwar-back/controllers"
)

func SetupRouter(router *gin.Engine) {
	router.GET("/api", controllers.SayHello)
}