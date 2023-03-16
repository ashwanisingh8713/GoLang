package products

import (
	"E-Commerce/database"
	"E-Commerce/database/table/products"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// CreateProduct godoc
// @Summary Creates Product.
// @Description To create and save Product Data.
// @Tags Product
// @Accept json
// @Produce json
// @Param       json  body BodyProductInput true "All fields are required"
// @Success 200 {object} userdata.StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /c_product [post]
func CreateProduct(c *gin.Context) {
	var in BodyProductInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if in.Price < 1 {
		c.JSON(http.StatusOK, gin.H{"status": false, "msg": "Price cannot be less than 1."})
		return
	}

	if in.Units <= 0 {
		c.JSON(http.StatusOK, gin.H{"status": false, "msg": "Unit cannot be less than 1."})
		return
	}

	if in.Weight <= 0 {
		c.JSON(http.StatusOK, gin.H{"status": false, "msg": "Weight cannot be less than 1 gram."})
		return
	}

	if len(strings.Trim(in.Name, " ")) == 0 {
		c.JSON(http.StatusOK, gin.H{"status": false, "msg": "Name cannot be empty"})
		return
	}

	status := products.Insert(database.DbInstance, in.CategoryId, in.Dimension, in.ManufacturerId, in.Name,
		in.Picture, in.Price, in.Description, in.SellerId, in.SellerName, in.Sku, in.SubCategoryId, in.Weight,
		in.Units, in.Subscribable)
	if status {
		c.JSON(http.StatusOK, gin.H{"status": status, "msg": "Product " + in.Name + " is created successfully."})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": status, "msg": "Failed to create Product " + in.Name + "."})
	}

}

// GetProduct godoc
// @Summary Get Product .
// @Description To get Product Data.
// @Tags Product
// @Accept json
// @Produce json
// @Param       json  body BodyProductIdInput true "It takes productId."
// @Success 200 {object} products.Product
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /r_product [post]
func GetProduct(c *gin.Context) {
	var in BodyProductIdInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := products.Read(database.DbInstance, in.ProductId)
	if product.ProductId.String() == in.ProductId {
		c.JSON(http.StatusOK, gin.H{"status": true, "data": product})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": false, "msg": "Failed to get Product."})
	}

}
