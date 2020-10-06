package controllers

import (
	"github.com/edwinvautier/DB_VAUTIER_P01/services"
	"github.com/edwinvautier/DB_VAUTIER_P01/models"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func GetEmployeesByOfficeCode(c *gin.Context) {
	code := services.ConvertStringToInt(c.Param("code"))
	var employees []models.Employee
	var err error

	employees, err = models.FindEmployeesByOfficeCode(code)
	fmt.Println(err)

	if err != nil {
		c.JSON(http.StatusNotFound, "Could'nt fetch office employees.")
		return
	}
	c.JSON(http.StatusOK, employees)
}

func GetEmployees(c *gin.Context) {
	var employees []models.EmployeeList
	var err error

	employees, err = models.FindEmployees()
	fmt.Println(err)

	if err != nil {
		c.JSON(http.StatusNotFound, "Could'nt fetch employees.")
		return
	}
	c.JSON(http.StatusOK, employees)
}