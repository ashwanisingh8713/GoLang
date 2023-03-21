package db_apis

import (
	"E-Commerce/database"
	"E-Commerce/database/table/address"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserAddress godoc
// @Summary Get all addresses of the User.
// @Description To get address list.
// @Tags Address
// @Accept json
// @Produce json
// @Param       json  body BodyUserIdInput true "It takes UserId"
// @Success 200 {object} []address.Address
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /getAddress [post]
func GetUserAddress(c *gin.Context) {
	var input BodyUserIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": address.Read(database.DbInstance, input.UserId)})
}

// CreateUserAddress godoc
// @Summary Creates or Add Address of the User.
// @Description Creates or Add Address of the User.
// @Tags Address
// @Accept json
// @Produce json
// @Param       json  body BodyAddressInput true "It takes UserId"
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /createAddress [post]
func CreateUserAddress(c *gin.Context) {
	var input BodyAddressInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var status, addressId = address.CreateAddress(database.DbInstance, input.UserId, input.AddressType, input.AddressLine1,
		input.AddressLine2, input.IsPreferred, input.Zip, input.City, input.State, input.Country, input.Mobile1, input.Mobile2)
	if status {
		c.JSON(http.StatusOK, gin.H{"status": status, "msg": "Successfully created new address.", "addressId": addressId})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": status, "msg": "Failed to create new address."})
	}
}

// DeleteAddress godoc
// @Summary Delete User's Address from the Address List.
// @Description Delete User's Address.
// @Tags Address
// @Accept json
// @Produce json
// @Param       json  body BodyActionOnIdInput true "It takes userId, addressId"
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /deleteAddress [delete]
func DeleteAddress(c *gin.Context) {
	var input BodyActionOnIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var status, addressId = address.Delete(database.DbInstance, input.UserId, input.ActionOnId)
	if status {
		c.JSON(http.StatusOK, gin.H{"status": status, "msg": "Successfully removed from your Address List.", "addressId": addressId})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": status, "msg": "Failed to remove from your Address List.", "addressId": addressId})
	}
}

// UpdateAddress godoc
// @Summary Update User's Address from the Address List.
// @Description Update User's Address.
// @Tags Address
// @Accept json
// @Produce json
// @Param       json  body BodyAddressUpdate true "All fields are required"
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /updateAddress [put]
func UpdateAddress(c *gin.Context) {
	var input BodyAddressUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var status, addressId = address.UpdateAddress(database.DbInstance, input.UserId, input.AddressId, input.AddressType, input.AddressLine1, input.AddressLine2,
		input.IsPreferred, input.Zip, input.City, input.State, input.Country, input.Mobile1, input.Mobile2)
	if status {
		c.JSON(http.StatusOK, gin.H{"status": status, "msg": "Successfully Updated Address.", "addressId": addressId})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": status, "msg": "Failed to update Address.", "addressId": addressId})
	}
}
