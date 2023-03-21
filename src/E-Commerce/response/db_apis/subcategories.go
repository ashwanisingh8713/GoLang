package db_apis

import (
	"E-Commerce/database"
	"E-Commerce/database/table/sub_category"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateSubCategory godoc
// @Summary Creates Subcategory.
// @Description To create Subcategory
// @Tags Subcategory
// @Accept json
// @Produce json
// @Param       json  body BodySubCategoryCreate true "It takes Manufacturer Name"
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /createSubCategory [post]
func CreateSubCategory(c *gin.Context) {
	var input BodySubCategoryCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var isSuccess, uniqueId = sub_category.CreateSubCategory(database.DbInstance, input.CategoryId, input.Name, input.Description)
	c.JSON(http.StatusOK, gin.H{"status": isSuccess, "msg": uniqueId})
}

// GetAllSubCategory godoc
// @Summary GetAll Subcategory.
// @Description To GetAll Subcategory
// @Tags Subcategory
// @Accept json
// @Produce json
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /getAllSubCategory [post]
func GetAllSubCategory(c *gin.Context) {
	var isSuccess, manufacturers = sub_category.GetAllSubCategory(database.DbInstance)
	c.JSON(http.StatusOK, gin.H{"status": isSuccess, "data": manufacturers})
}

// GetAllSubCategoryByCateId godoc
// @Summary Get Subcategory By ID.
// @Description To Get Subcategory By ID
// @Tags Subcategory
// @Accept json
// @Produce json
// @Param       json  body BodyById true "It takes Category id"
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /getAllSubCategoryByCateId [post]
func GetAllSubCategoryByCateId(c *gin.Context) {
	var input BodyById
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var isSuccess, manufacturer = sub_category.GetSubCategoryById(database.DbInstance, input.ID)
	c.JSON(http.StatusOK, gin.H{"status": isSuccess, "data": manufacturer})
}
