package definitions

import (
	"github.com/binodluitel/api/pkg/models/users"
	"github.com/gin-gonic/gin"
)

//go:generate ../../../.build/bin/mockery --name=UsersService

// UsersService defines a methods for users REST API service
type UsersService interface {
	CreateUser(ctx *gin.Context, request *users.CreateRequest) (*users.User, error)
	GetUser(ctx *gin.Context, id, filters string) (*users.User, error)
	ListUsers(ctx *gin.Context, filters string) ([]*users.User, error)
	UpdateUser(ctx *gin.Context, id string, request *users.UpdateRequest) (*users.User, error)
	DeleteUser(ctx *gin.Context, id string) (*users.User, error)
}
