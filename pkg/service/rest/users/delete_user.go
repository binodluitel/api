package users

import (
	"github.com/binodluitel/api/pkg/log"
	"github.com/binodluitel/api/pkg/models/users"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (*Users) DeleteUser(ctx *gin.Context, id string) (*users.User, error) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	logger.Debug("deleting user", zap.String("id", id))
	return &users.User{}, nil
}
