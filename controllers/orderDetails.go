package controllers

import (
	"fmt"
	"net/http"

	"github.com/edwinvautier/DB_VAUTIER_P01/models"
	"github.com/edwinvautier/DB_VAUTIER_P01/services"
	"github.com/gin-gonic/gin"
)

func GetOrderDetailsByOrderNumber(c *gin.Context) {
	number := services.ConvertStringToInt(c.Param("number"))
	var details []models.OrderDetails
	var err error

	details, err = models.FindOrderDetailsByOrderNumber(number)
	fmt.Println(err)

	if err != nil {
		c.JSON(http.StatusNotFound, "Could'nt fetch order details.")
		return
	}
	c.JSON(http.StatusOK, details)
}
