package users

import (
	"github.com/binodluitel/api/pkg/log"
	"github.com/binodluitel/api/pkg/models/users"
	"github.com/gin-gonic/gin"
)

func (*Users) CreateUser(ctx *gin.Context, request *users.CreateRequest) (*users.User, error) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	logger.Debug("creating user")
	return &users.User{}, nil
}
