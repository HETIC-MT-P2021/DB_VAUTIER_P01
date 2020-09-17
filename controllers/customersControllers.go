package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/edwinvautier/DB_VAUTIER_P01/services"
	"github.com/edwinvautier/DB_VAUTIER_P01/models"
	"net/http"
)

func GetCustomerByNumber(c *gin.Context) {
	number := services.ConvertStringToInt(c.Param("number"))
	var customer models.Customer
	var err error

	customer, err = models.FindCustomerByNumber(number)
	if err != nil {
		c.JSON(http.StatusNotFound, "Could'nt fetch user.")
	}
	c.JSON(http.StatusOK, customer)
}