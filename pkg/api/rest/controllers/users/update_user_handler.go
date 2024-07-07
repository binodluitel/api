package users

import (
	"net/http"

	"github.com/binodluitel/api/pkg/log"
	"github.com/binodluitel/api/pkg/models/users"
	"github.com/gin-gonic/gin"
)

func (c *Controller) UpdateUser(ctx *gin.Context) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	userID := ctx.Param("user_id")
	request := &users.UpdateRequest{}
	user, err := c.service.UpdateUser(ctx, userID, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
