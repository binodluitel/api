package stream

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/binodluitel/api/pkg/log"
	"github.com/binodluitel/api/pkg/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"k8s.io/utils/ptr"
)

func (c *Controller) StreamLogs(ctx *gin.Context) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	logger.Info("streaming logs")
	ctx.Stream(func(w io.Writer) bool {
		request := new(models.StreamRequest)
		c.setRequestDefaults(ctx, request)
		for {
			select {
			case <-ctx.Done():
				return false
			default:
				logs, err := c.service.StreamLogs(ctx, request)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return false
				}
				defer logs.Close()
				if _, err := io.Copy(w, logs); err != nil {
					// If the client disconnects, stop writing logs
					if netErr, ok := err.(*net.OpError); ok && netErr.Err.Error() == "write: broken pipe" {
						return false
					}
					logger.Error("error streaming logs", zap.Error(err))
					return false
				}
				if !request.Follow {
					return false
				}

				// Update the request with timestamp of the last log line
				var buf bytes.Buffer
				_, err = io.Copy(&buf, logs)
				if err != nil {
					logger.Error("error copying logs to buffer", zap.Error(err))
					return false
				}
				if timestamp := lastTimestamp(ctx, &buf); timestamp != "" {
					since, err := time.Parse(time.RFC3339Nano, timestamp)
					if err != nil {
						logger.Error("error parsing timestamp", zap.Error(err))
						return false
					}
					request.SinceSeconds = ptr.To[int64](int64(time.Since(since).Seconds()))
				}
			}
		}
	})
}

func (c *Controller) setRequestDefaults(ctx *gin.Context, request *models.StreamRequest) {
	if request == nil {
		request = new(models.StreamRequest)
	}
	request.Follow = ctx.Query("follow") == "true"
	if sinceSeconds := ctx.Query("since_seconds"); sinceSeconds != "" {
		since, err := strconv.ParseInt(sinceSeconds, 10, 64)
		if err == nil {
			request.SinceSeconds = &since
		}
	}
	if request.SinceSeconds == nil {
		request.SinceSeconds = ptr.To[int64](int64(time.Duration(48 * time.Hour).Seconds()))
	}
	request.Container = ctx.Query("container")
	if request.Container == "" {
		request.Container = "api"
	}
	request.Namespace = ctx.Query("namespace")
	if request.Namespace == "" {
		request.Namespace = "api"
	}
}

func lastTimestamp(ctx *gin.Context, logs io.Reader) string {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	r := bufio.NewReader(logs)
	var timestamp string
	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				logger.Error("Error reading line:", zap.Error(err))
				return ""
			}
		}
		if len(line) == 0 {
			continue
		}
		if idx := strings.IndexRune(strings.TrimSuffix(string(line), "\n"), ' '); idx != -1 {
			timestamp = string(line[:idx])
		}
	}
	logger.Info("last timestamp", zap.String("timestamp", timestamp))
	return timestamp
}
