package db_apis

import (
	"E-Commerce/database"
	"E-Commerce/database/table/category"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateCategory godoc
// @Summary Creates Category.
// @Description To create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param       json  body BodySubCategoryCreate true "It takes Manufacturer Name"
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /createCategory [post]
func CreateCategory(c *gin.Context) {
	var input BodySubCategoryCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var isSuccess, uniqueId = category.CreateCategory(database.DbInstance, input.CategoryId, input.Name, input.Description)
	c.JSON(http.StatusOK, gin.H{"status": isSuccess, "msg": uniqueId})
}

// GetAllCategory godoc
// @Summary GetAll Category.
// @Description To GetAll Category
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /getAllCategory [post]
func GetAllCategory(c *gin.Context) {
	var isSuccess, manufacturers = category.GetAllCategory(database.DbInstance)
	c.JSON(http.StatusOK, gin.H{"status": isSuccess, "data": manufacturers})
}

// GetAllCategoryById godoc
// @Summary Get Category By ID.
// @Description To Get Category By ID
// @Tags Category
// @Accept json
// @Produce json
// @Param       json  body BodyById true "It takes Category id"
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /getAllCategoryById [post]
func GetAllCategoryById(c *gin.Context) {
	var input BodyById
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var isSuccess, manufacturer = category.GetCategoryById(database.DbInstance, input.ID)
	c.JSON(http.StatusOK, gin.H{"status": isSuccess, "data": manufacturer})
}
