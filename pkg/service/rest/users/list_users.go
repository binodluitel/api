package users

import (
	"github.com/binodluitel/api/pkg/log"
	"github.com/binodluitel/api/pkg/models/users"
	"github.com/gin-gonic/gin"
)

func (*Users) ListUsers(ctx *gin.Context, filters string) ([]*users.User, error) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	logger.Debug("listing users")
	return []*users.User{}, nil
}
