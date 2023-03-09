package userData

import (
	"E-Commerce/database"
	"E-Commerce/database/table/address"
	cart_item2 "E-Commerce/database/table/cart_item"
	"E-Commerce/database/table/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserIdInput struct {
	UserId string `json:"userId" binding:"required"`
}

// GetUserInfo /**
func GetUserInfo(c *gin.Context) {
	var input CreateUserIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user = user.Read(database.DbInstance, input.UserId)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// GetUserAddress  /**
func GetUserAddress(c *gin.Context) {
	var input CreateUserIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var address = address.Read(database.DbInstance, input.UserId)
	c.JSON(http.StatusOK, gin.H{"data": address})
}

// GetUserCartItem  /**
func GetUserCartItem(c *gin.Context) {
	var input CreateUserIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var cartItems = cart_item2.Read(database.DbInstance, input.UserId)
	c.JSON(http.StatusOK, gin.H{"data": cartItems})
}
