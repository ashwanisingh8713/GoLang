package userData

import (
	"E-Commerce/database"
	"E-Commerce/database/table/address"
	"E-Commerce/database/table/cart_item"
	"E-Commerce/database/table/subscriptions"
	"E-Commerce/database/table/user"
	"E-Commerce/database/table/wishlist"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateWishlistInput struct {
	UserId    string `json:"userId" binding:"required"`
	ProductId string `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

// GetUserInfo /**
func GetUserInfo(c *gin.Context) {
	var input CreateUserIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user.Read(database.DbInstance, input.UserId)})
}

// GetUserAddress  /**
func GetUserAddress(c *gin.Context) {
	var input CreateUserIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": address.Read(database.DbInstance, input.UserId)})
}

// GetUserCartItem  /**
func GetUserCartItem(c *gin.Context) {
	var input CreateUserIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var cartItems = cart_item.Read(database.DbInstance, input.UserId)
	c.JSON(http.StatusOK, gin.H{"data": cartItems})
}

// GetUserAllSubscription  /**
func GetUserAllSubscription(c *gin.Context) {
	var input CreateUserIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var subscriptions = subscriptions.ReadAllSubcriptions(database.DbInstance, input.UserId)
	c.JSON(http.StatusOK, gin.H{"data": subscriptions})
}

// CreateWishlist  /**
func CreateWishlist(c *gin.Context) {
	var input CreateWishlistInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var isSuccess, msg = wishlist.Insert(database.DbInstance, input.UserId, input.ProductId, input.Quantity)
	c.JSON(http.StatusOK, gin.H{"data": isSuccess, "msg": msg})
}

// swagger:parameters createUserIdInput
type CreateUserIdInput struct {
	UserId string `json:"userId" binding:"required"`
}

// GetWishlist godoc
// @Summary Shows the all wishlist of the User.
// @Description Get the all wishlist of the User.
// @Tags User
// @Accept json
// @Produce json
// @Param       json  body CreateUserIdInput true "Request Body to get all Wishlist"
// @Success 200 {object} []wishlist.Wishlist
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /get_wishlist [post]
func GetWishlist(c *gin.Context) {
	var input CreateUserIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wishlist.Read(database.DbInstance, input.UserId)})
}

// DeleteWishlist godoc
// @Summary Deletes the selected wishlistId from the Wishlist Table.
// @Description To delete the selected product from User's Wishlist.
// @Tags User
// @Accept json
// @Produce json
// @Param       wishlistId  query string true "wishlist_id value"
// @Param       wishlistId  query string true "wishlist_id value"
// @Success 200 {object} []wishlist.Wishlist
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /wishlist [delete]
func DeleteWishlist(c *gin.Context) {
	var wishlistId = c.Query("wishlistId")
	var isSuccess, msg = wishlist.Delete(database.DbInstance, wishlistId)
	c.JSON(http.StatusOK, gin.H{"data": isSuccess, "msg": msg})
}
