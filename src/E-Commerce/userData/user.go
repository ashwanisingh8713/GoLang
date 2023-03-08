package userData

import (
	"E-Commerce/database"
	"E-Commerce/database/table/address"
	"E-Commerce/database/table/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	user.Read(database.DbInstance)
	var user = user.Read(database.DbInstance)
	var address = address.Read(database.DbInstance, user.UserId)
	c.JSON(http.StatusOK, gin.H{"user": user, "address": address})
	//c.JSON(http.StatusOK, gin.H{"Ashwani": "Singh"})
}
