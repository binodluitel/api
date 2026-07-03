package readiness

import (
	"context"
	"net/http"
	"time"

	"github.com/binodluitel/api/pkg/log"
	"github.com/gin-gonic/gin"
)

func (c *Controller) ReadyZ(ctx *gin.Context) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	readyCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	if restChecker := c.service; restChecker != nil {
		if checker, ok := restChecker.(interface{ Ready(context.Context) error }); ok {
			if err := checker.Ready(readyCtx); err != nil {
				ctx.JSON(http.StatusServiceUnavailable, gin.H{
					"ready": false,
					"error": err.Error(),
				})
				return
			}
		} else {
			ctx.JSON(http.StatusServiceUnavailable, gin.H{
				"ready": false,
				"error": "pods service does not support readiness checks",
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"ready": true})
}
