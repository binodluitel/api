package users

import (
	"net/http"

	"github.com/binodluitel/api/pkg/log"
	"github.com/binodluitel/api/pkg/models/users"
	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateUser(ctx *gin.Context) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	request := &users.CreateRequest{}
	user, err := c.service.CreateUser(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
