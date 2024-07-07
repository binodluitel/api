package users

import (
	svcdef "github.com/binodluitel/api/pkg/service/definitions"
	"github.com/gin-gonic/gin"
)

// Controller is user REST API controller
type Controller struct {
	service svcdef.UsersService
}

func New(service svcdef.UsersService, router *gin.RouterGroup) *Controller {
	c := &Controller{service: service}
	usersRouter := router.Group("")
	usersRouter.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	})
	usersRouter.GET("/users/:user_id", c.GetUser)
	usersRouter.GET("/users", c.ListUsers)
	usersRouter.POST("/users", c.CreateUser)
	usersRouter.PATCH("/users/:user_id", c.UpdateUser)
	usersRouter.DELETE("/users/:user_id", c.DeleteUser)
	return c
}
