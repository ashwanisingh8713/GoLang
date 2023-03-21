package db_apis

import (
	"E-Commerce/database"
	"E-Commerce/database/table/manufacturer"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateManufacturer godoc
// @Summary Creates Manufacturer.
// @Description To create Manufacturer
// @Tags Manufacturer
// @Accept json
// @Produce json
// @Param       json  body BodyNameCreate true "It takes Manufacturer Name"
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /createManufacturer [post]
func CreateManufacturer(c *gin.Context) {
	var input BodyNameCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var isSuccess, uniqueId = manufacturer.CreateManufacturer(database.DbInstance, input.Name)
	c.JSON(http.StatusOK, gin.H{"status": isSuccess, "msg": uniqueId})
}

// GetAllManufacturer godoc
// @Summary GetAll Manufacturer.
// @Description To GetAll Manufacturer
// @Tags Manufacturer
// @Accept json
// @Produce json
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /getAllManufacturer [post]
func GetAllManufacturer(c *gin.Context) {
	var isSuccess, manufacturers = manufacturer.GetAllManufacturer(database.DbInstance)
	c.JSON(http.StatusOK, gin.H{"status": isSuccess, "data": manufacturers})
}

// GetManufacturerById godoc
// @Summary Get Manufacturer By ID.
// @Description To Get Manufacturer By ID
// @Tags Manufacturer
// @Accept json
// @Produce json
// @Param       json  body BodyById true "It takes Manufacturer  Id"
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /getManufacturerById [post]
func GetManufacturerById(c *gin.Context) {
	var input BodyById
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var isSuccess, manufacturer = manufacturer.GetManufacturerById(database.DbInstance, input.ID)
	c.JSON(http.StatusOK, gin.H{"status": isSuccess, "data": manufacturer})
}
