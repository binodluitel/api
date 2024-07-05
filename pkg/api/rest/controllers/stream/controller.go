package stream

import (
	"net/http"

	svcdef "github.com/binodluitel/api/pkg/service/definitions"
	"github.com/gin-gonic/gin"
)

// Controller is user REST API controller
type Controller struct {
	service svcdef.StreamService
}

func New(service svcdef.StreamService, router *gin.RouterGroup) *Controller {
	c := &Controller{service: service}
	router.Any("/stream", func(c *gin.Context) { c.Status(http.StatusTeapot) })
	router.GET("/stream/logs", c.StreamLogs)
	return c
}
