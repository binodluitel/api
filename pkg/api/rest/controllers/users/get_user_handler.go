package users

import (
	"net/http"

	"github.com/binodluitel/api/pkg/log"
	"github.com/gin-gonic/gin"
)

func (c *Controller) GetUser(ctx *gin.Context) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	userID := ctx.Param("user_id")
	filters := ""
	user, err := c.service.GetUser(ctx, userID, filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
