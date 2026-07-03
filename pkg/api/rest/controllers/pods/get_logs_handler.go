package pods

import (
	"io"
	"net/http"

	"github.com/binodluitel/api/pkg/log"
	"github.com/binodluitel/api/pkg/models/pods"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	// streamBufferSize is the size of the buffer used to stream logs
	streamBufferSize = 8192
)

func (c *Controller) GetLogs(ctx *gin.Context) {
	_, logger := log.Get(ctx)
	request := new(pods.Logs)
	request.PodName = ctx.Param("pod_name")
	if request.PodName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "pod name is required"})
		return
	}
	logger.Debug("Getting pod logs", zap.String("pod_name", request.PodName))
	c.setRequestDefaults(ctx, request)
	logs, err := c.service.GetLogs(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func() { _ = logs.Close() }()
	ctx.Header("Content-Type", "text/plain")
	ctx.Header("Transfer-Encoding", "chunked")

	buf := make([]byte, streamBufferSize)
	for reading := true; reading; {
		reading = ctx.Stream(func(w io.Writer) bool {
			select {
			case <-ctx.Done():
				logger.Info("client is disconnected")
				return false
			default:
				n, err := logs.Read(buf)
				if n > 0 {
					if _, errWrite := w.Write(buf[:n]); errWrite != nil {
						logger.Error("failed to write logs to client", zap.Error(errWrite))
						return false
					}
					return true
				}
				if err != nil && err != io.EOF {
					logger.Error("failed to read logs", zap.Error(err))
				}
			}
			return false
		})
	}
}

func (c *Controller) setRequestDefaults(ctx *gin.Context, request *pods.Logs) {
	if request == nil {
		request = new(pods.Logs)
	}
	request.Follow = ctx.Query("follow") == "true"
	request.Container = ctx.Query("container")
	if request.Container == "" {
		request.Container = "api"
	}
	request.Namespace = ctx.Query("namespace")
	if request.Namespace == "" {
		request.Namespace = "api"
	}
}
