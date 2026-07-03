package readiness

import (
	svcdef "github.com/binodluitel/api/pkg/service/definitions"
	"github.com/gin-gonic/gin"
)

// Controller is readiness REST API controller
type Controller struct {
	service svcdef.PodsService
}

func New(service svcdef.PodsService, router *gin.RouterGroup) *Controller {
	c := &Controller{service: service}
	readinessRouter := router.Group("")
	readinessRouter.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Next()
	})

	// Readiness check endpoints for Kubernetes
	router.GET("/readyz", c.ReadyZ)
	return c
}
