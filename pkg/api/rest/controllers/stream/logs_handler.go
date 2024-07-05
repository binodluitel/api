package stream

import (
	"io"
	"net"

	"github.com/binodluitel/api/pkg/log"
	"github.com/binodluitel/api/pkg/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) StreamLogs(ctx *gin.Context) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	ctx.Stream(func(w io.Writer) bool {
		for {
			select {
			case <-ctx.Done():
				return false
			default:
				request := new(models.StreamRequest)
				logs, err := c.service.StreamLogs(ctx, request)
				if err != nil {
					ctx.JSON(500, gin.H{"error": err.Error()})
					return false
				}
				if logs != nil {
					logger.Info("writing logs", zap.String("logs", *logs))
					_, err = w.Write([]byte(*logs))
					if err != nil {
						// If the client disconnects, stop writing logs
						if netErr, ok := err.(*net.OpError); ok && netErr.Err.Error() == "write: broken pipe" {
							return false
						}
						logger.Error("error writing logs", zap.Error(err))
						return false
					}
				}
			}
		}
	})
}
