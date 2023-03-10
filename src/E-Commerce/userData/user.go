package userData

import (
	"E-Commerce/database"
	"E-Commerce/database/table/address"
	"E-Commerce/database/table/cart_item"
	"E-Commerce/database/table/subscriptions"
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

	c.JSON(http.StatusOK, gin.H{"user": user.Read(database.DbInstance, input.UserId)})
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
