package users

import (
	"net/http"

	"github.com/binodluitel/api/pkg/log"
	"github.com/gin-gonic/gin"
)

func (c *Controller) DeleteUser(ctx *gin.Context) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	userID := ctx.Param("user_id")
	user, err := c.service.DeleteUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, user) // returning 202 Accepted status code since the deletion is async
}
