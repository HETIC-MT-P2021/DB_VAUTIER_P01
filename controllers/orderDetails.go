package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/edwinvautier/DB_VAUTIER_P01/services"
	"github.com/edwinvautier/DB_VAUTIER_P01/models"
	"net/http"
	"fmt"
)

func GetOrdersDetailsByOrderNumber(c * gin.Context) {
	number := services.ConvertStringToInt(c.Param("number"))
	var details []models.OrderDetail
	var err error

	details, err = models.FindOrderDetailsByOrderNumber(number)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusNotFound, "Could'nt fetch order details.")
		return
	}
	c.JSON(http.StatusOK, details)
}