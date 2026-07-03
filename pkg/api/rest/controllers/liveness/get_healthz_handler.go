package liveness

import (
	"net/http"

	"github.com/binodluitel/api/pkg/log"
	"github.com/gin-gonic/gin"
)

const statusAlive = "alive"

func (c *Controller) HealthZ(ctx *gin.Context) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	ctx.JSON(http.StatusOK, gin.H{"status": statusAlive})
}
