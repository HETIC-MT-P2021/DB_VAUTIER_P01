package controllers

import (
	"fmt"
	"net/http"

	"github.com/edwinvautier/DB_VAUTIER_P01/models"
	"github.com/edwinvautier/DB_VAUTIER_P01/services"
	"github.com/gin-gonic/gin"
)

func GetOfficeByCode(c *gin.Context) {
	code := services.ConvertStringToInt(c.Param("code"))
	var office models.Office
	var err error

	office, err = models.FindOfficeByCode(code)
	fmt.Println(err)

	if err != nil {
		c.JSON(http.StatusNotFound, "Could'nt fetch office.")
		return
	}
	c.JSON(http.StatusOK, office)
}
