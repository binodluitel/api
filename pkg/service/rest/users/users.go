package users

import (
	"github.com/binodluitel/api/pkg/config"
	svcdef "github.com/binodluitel/api/pkg/service/definitions"
)

// Users defines users service instance
type Users struct{}

// New creates and returns a new user service instance
func New(cfg *config.Config) (svcdef.UsersService, error) {
	return &Users{}, nil
}
