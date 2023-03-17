package userdata

import (
	"E-Commerce/database"
	"E-Commerce/database/table/subscriptions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserAllSubscription godoc
// @Summary Get all subscribed products of the User.
// @Description To get subscription list.
// @Tags Subscription
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
