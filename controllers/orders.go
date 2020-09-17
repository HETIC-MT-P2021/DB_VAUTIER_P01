package controllers

import (
	"fmt"
	"net/http"

	"github.com/edwinvautier/DB_VAUTIER_P01/models"
	"github.com/edwinvautier/DB_VAUTIER_P01/services"
	"github.com/gin-gonic/gin"
)

func GetOrdersByCustomerNumber(c *gin.Context) {
	number := services.ConvertStringToInt(c.Param("number"))
	var orders []models.Order
	var err error

	orders, err = models.FindOrdersByCustomerNumber(number)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusNotFound, "Could'nt fetch orders.")
		return
	}
	c.JSON(http.StatusOK, orders)
}

func GetOrderByNumber(c *gin.Context) {
	number := services.ConvertStringToInt(c.Param("number"))
	var order models.OrderFull
	var err error

	order, err = models.FindOrderByNumber(number)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusNotFound, "Could'nt fetch order.")
		return
	}
	c.JSON(http.StatusOK, order)
}