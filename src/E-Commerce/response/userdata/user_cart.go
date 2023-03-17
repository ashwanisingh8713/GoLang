package userdata

import (
	"E-Commerce/database"
	"E-Commerce/database/table/cart_item"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCartItem godoc
// @Summary Get all cart products of the User.
// @Description To get cart list.
// @Tags Cart
// @Accept json
// @Produce json
// @Param       json  body BodyUserIdInput true "It takes UserId"
// @Success 200 {object} []cart_item.CartItem
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /getCart [post]
func GetCartItem(c *gin.Context) {
	var input BodyUserIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var cartItems = cart_item.Read(database.DbInstance, input.UserId)
	c.JSON(http.StatusOK, gin.H{"data": cartItems})
}

// CreateCartItem godoc
// @Summary Add item into the Cart of the User.
// @Description Create or Add Cart Item
// @Tags Cart
// @Accept json
// @Produce json
// @Param       json  body BodySelectedProductInput true "It takes UserId, productId, quantity"
// @Success 200 {object} []cart_item.CartItem
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /createCart [post]
func CreateCartItem(c *gin.Context) {
	var input BodySelectedProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Quantity < 1 {
		c.JSON(http.StatusOK, gin.H{"status": false, "msg": "Quantity cannot be less than 1."})
		return
	}
	var status = cart_item.Insert(database.DbInstance, input.UserId, input.ProductId, input.Quantity)
	if status {
		c.JSON(http.StatusOK, gin.H{"status": status, "msg": "Successfully added in your cart."})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": status, "msg": "Failed to add in your cart."})
	}
}

// DeleteCartItem godoc
// @Summary Delete cart item from the Cart.
// @Description Delete cart item from the Cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param       json  body BodyActionOnIdInput true "It takes userId, productId"
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /deleteCart [delete]
func DeleteCartItem(c *gin.Context) {
	var input BodyActionOnIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var status, cartItemId = cart_item.Delete(database.DbInstance, input.UserId, input.ActionOnId)
	if status {
		c.JSON(http.StatusOK, gin.H{"status": status, "msg": "Successfully removed from your cart.", "productId": cartItemId})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": status, "msg": "Failed to remove from your cart.", "productId": cartItemId})
	}
}
