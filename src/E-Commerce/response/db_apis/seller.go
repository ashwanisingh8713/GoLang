package db_apis

import (
	"E-Commerce/database"
	"E-Commerce/database/table/seller"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateSeller godoc
// @Summary Creates Seller.
// @Description To create Seller
// @Tags Seller
// @Accept json
// @Produce json
// @Param       json  body BodyNameCreate true "It takes Manufacturer Name"
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /createSeller [post]
func CreateSeller(c *gin.Context) {
	var input BodyNameCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var isSuccess, uniqueId = seller.CreateSeller(database.DbInstance, input.Name)
	c.JSON(http.StatusOK, gin.H{"status": isSuccess, "msg": uniqueId})
}

// GetAllSeller godoc
// @Summary GetAll Seller.
// @Description To GetAll Seller
// @Tags Seller
// @Accept json
// @Produce json
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /getAllSeller [post]
func GetAllSeller(c *gin.Context) {
	var isSuccess, manufacturers = seller.GetAllSeller(database.DbInstance)
	c.JSON(http.StatusOK, gin.H{"status": isSuccess, "data": manufacturers})
}

// GetSellerById godoc
// @Summary Get Seller By ID.
// @Description To Get Seller By ID
// @Tags Seller
// @Accept json
// @Produce json
// @Param       json  body BodyById true "It takes Seller id"
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /getSellerById [post]
func GetSellerById(c *gin.Context) {
	var input BodyById
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var isSuccess, manufacturer = seller.GetSellerById(database.DbInstance, input.ID)
	c.JSON(http.StatusOK, gin.H{"status": isSuccess, "data": manufacturer})
}
