package liveness

import (
	"github.com/gin-gonic/gin"
)

// Controller is liveness REST API controller
type Controller struct{}

func New(router *gin.RouterGroup) *Controller {
	c := &Controller{}
	livenessRouter := router.Group("")
	livenessRouter.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Next()
	})

	// Liveness check endpoints for Kubernetes
	router.GET("/healthz", c.HealthZ)
	return c
}
