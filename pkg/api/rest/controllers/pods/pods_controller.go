package pods

import (
	"net/http"

	svcdef "github.com/binodluitel/api/pkg/service/definitions"
	"github.com/gin-gonic/gin"
)

// Controller is user REST API controller
type Controller struct {
	service svcdef.PodsService
}

func New(service svcdef.PodsService, router *gin.RouterGroup) *Controller {
	c := &Controller{service: service}
	podsRouter := router.Group("")
	podsRouter.Any("/pods", func(c *gin.Context) { c.Status(http.StatusTeapot) })
	podsRouter.GET("/pods/:pod_name/logs", c.GetLogs)
	return c
}
