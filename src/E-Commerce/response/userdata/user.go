package userdata

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

// CreateUser godoc
// @Summary Creates User .
// @Description To create and save User Data.
// @Tags USER
// @Accept json
// @Produce json
// @Param       json  body BodyCreateUserInput true "It takes UserId"
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /createUser [post]
func CreateUser(c *gin.Context) {
	var input BodyCreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var status = user.CreateUser(database.DbInstance, input.FirstName, input.LastName, input.Mobile, input.Email, input.Password, input.LoginMode)
	if status {
		c.JSON(http.StatusOK, gin.H{"status": status, "msg": "User " + input.FirstName + " is created successfully."})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": status, "msg": "Failed to create User " + input.FirstName + "."})
	}

}

// GetAllUsers godoc
// @Summary Provides All Users.
// @Description To get all Users list (userId = 09876).
// @Tags USER
// @Accept json
// @Produce json
// @Param       json  body BodyUserIdInput true "It takes UserId"
// @Success 200 {object} []user.User
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /getAllUsers [post]
func GetAllUsers(c *gin.Context) {
	var input BodyUserIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if this request is coming from valid user
	if input.UserId == "09876" {
		var users, status = user.GetAllUser(database.DbInstance)
		c.JSON(http.StatusOK, gin.H{"status": status, "users": users})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": false, "msg": "To get all Users info you need the admin UserId"})
	}

}

// GetUserInfo godoc
// @Summary Get User's information.
// @Description To get User's information.
// @Tags USER
// @Accept json
// @Produce json
// @Param       json  body BodyUserIdInput true "It takes UserId"
// @Success 200 {object} user.User
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /user [post]
func GetUserInfo(c *gin.Context) {
	var input BodyUserIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user.GetUserInfo(database.DbInstance, input.UserId)})
}

// GetUserAddress godoc
// @Summary Get all addresses of the User.
// @Description To get address list.
// @Tags USER
// @Accept json
// @Produce json
// @Param       json  body BodyUserIdInput true "It takes UserId"
// @Success 200 {object} []address.Address
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /address [post]
func GetUserAddress(c *gin.Context) {
	var input BodyUserIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": address.Read(database.DbInstance, input.UserId)})
}

// GetUserCartItem godoc
// @Summary Get all cart products of the User.
// @Description To get cart list.
// @Tags USER_ITEM
// @Accept json
// @Produce json
// @Param       json  body BodyUserIdInput true "It takes UserId"
// @Success 200 {object} []cart_item.CartItem
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /cart_item [post]
func GetUserCartItem(c *gin.Context) {
	var input BodyUserIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var cartItems = cart_item.Read(database.DbInstance, input.UserId)
	c.JSON(http.StatusOK, gin.H{"data": cartItems})
}

// GetUserAllSubscription godoc
// @Summary Get all subscribed products of the User.
// @Description To get subscription list.
// @Tags USER_ITEM
// @Accept json
// @Produce json
// @Param       json  body BodyUserIdInput true "It takes UserId"
// @Success 200 {object} []subscriptions.Subscription
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /subscription [post]
func GetUserAllSubscription(c *gin.Context) {
	var input BodyUserIdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var subscriptions = subscriptions.ReadAllSubcriptions(database.DbInstance, input.UserId)
	c.JSON(http.StatusOK, gin.H{"data": subscriptions})
}

// CreateWishlist godoc
// @Summary Creates wishlist for selected product.
// @Description To create wishlist of the User
// @Tags USER_ITEM
// @Accept json
// @Produce json
// @Param       json  body BodyWishlistInput true "It takes UserId, ProductId, Quantity"
// @Success 200 {object} StatusMsg
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /c_wishlist [post]
func CreateWishlist(c *gin.Context) {
	var input BodyWishlistInput
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
// @Tags USER_ITEM
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
// @Tags USER_ITEM
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
