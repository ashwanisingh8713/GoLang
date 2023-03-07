package userData

import (
	"ServerDrivenUI/src/E-Commerce/database"
	"ServerDrivenUI/src/cassandra/db/table/address"
	"ServerDrivenUI/src/cassandra/db/table/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	user.Read(database.DbInstance)
	var user = user.Read(database.DbInstance)
	var address = address.Read(database.DbInstance, user.UserId)
	c.JSON(http.StatusOK, gin.H{"user": user, "address": address})
}