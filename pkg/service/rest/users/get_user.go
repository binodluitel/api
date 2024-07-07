package users

import (
	"github.com/binodluitel/api/pkg/log"
	"github.com/binodluitel/api/pkg/models/users"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (*Users) GetUser(ctx *gin.Context, id, filters string) (*users.User, error) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	logger.Debug("getting user", zap.String("id", id))
	return &users.User{}, nil
}
