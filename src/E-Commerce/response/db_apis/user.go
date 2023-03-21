package db_apis

import (
	"E-Commerce/database"
	"E-Commerce/database/table/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser godoc
// @Summary Creates User .
// @Description To create and save User Data.
// @Tags User
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
	var status, uniqueUserId = user.CreateUser(database.DbInstance, input.FirstName, input.LastName, input.Mobile, input.Email, input.Password, input.LoginMode)
	if status {
		c.JSON(http.StatusOK, gin.H{"status": status, "msg": "User " + input.FirstName + " is created successfully.", "userId": uniqueUserId})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": status, "msg": "Failed to create User " + input.FirstName + "."})
	}

}

// GetAllUsers godoc
// @Summary Provides All Users.
// @Description To get all Users list (userId = 09876).
// @Tags User
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
// @Tags User
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

// GetUserLogin godoc
// @Summary User Login.
// @Description User Login using either email or mobile.
// @Tags User
// @Accept json
// @Produce json
// @Param       json  body BodyUserLoginInput true "It takes Email or Mobile and Password"
// @Success 200 {object} user.User
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /loginUser [post]
func GetUserLogin(c *gin.Context) {
	var input BodyUserLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var status, user = user.GetUserLogin(database.DbInstance, input.EmailOrMobile, input.Password)

	if status {
		c.JSON(http.StatusOK, gin.H{"data": user})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "kindly enter valid user credentials."})
	}
}
