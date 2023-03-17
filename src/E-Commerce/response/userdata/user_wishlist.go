package userdata

import (
	"E-Commerce/database"
	"E-Commerce/database/table/wishlist"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateWishlist godoc
// @Summary Creates wishlist for selected product.
// @Description To create wishlist of the User
// @Tags Wishlist
// @Accept json
// @Produce json
// @Param       json  body BodySelectedProductInput true "It takes UserId, ProductId, Quantity"
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /c_wishlist [post]
func CreateWishlist(c *gin.Context) {
	var input BodySelectedProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var isSuccess, msg = wishlist.Insert(database.DbInstance, input.UserId, input.ProductId, input.Quantity)
	c.JSON(http.StatusOK, gin.H{"status": isSuccess, "msg": msg})
}

// GetWishlist godoc
// @Summary Shows the all wishlist of the User.
// @Description Get the all wishlist of the User.
// @Tags Wishlist
// @Accept json
// @Produce json
// @Param       json  body BodyUserIdInput true "It takes UserId"
// @Success 200 {object} []wishlist.Wishlist
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /r_wishlist [post]
func GetWishlist(c *gin.Context) {
	var input BodyUserIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wishlist.Read(database.DbInstance, input.UserId)})
}

// DeleteWishlist godoc
// @Summary Deletes the selected wishlistId from the Wishlist Table.
// @Description To delete the selected product from User's Wishlist.
// @Tags Wishlist
// @Accept json
// @Produce json
// @Param       wishlistId  query string true "wishlist_id value"
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /d_wishlist [delete]
func DeleteWishlist(c *gin.Context) {
	var wishlistId = c.Query("wishlistId")
	var isSuccess, msg = wishlist.Delete(database.DbInstance, wishlistId)
	c.JSON(http.StatusOK, gin.H{"status": isSuccess, "msg": msg})
}
