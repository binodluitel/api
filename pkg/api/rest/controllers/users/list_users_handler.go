package users

import (
	"net/http"

	"github.com/binodluitel/api/pkg/log"
	"github.com/gin-gonic/gin"
)

func (c *Controller) ListUsers(ctx *gin.Context) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	filters := ""
	users, err := c.service.ListUsers(ctx, filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
