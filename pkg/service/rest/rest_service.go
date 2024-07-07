package rest

import (
	"fmt"

	"github.com/binodluitel/api/pkg/config"
	svcdef "github.com/binodluitel/api/pkg/service/definitions"
	podssvc "github.com/binodluitel/api/pkg/service/rest/pods"
	userssvc "github.com/binodluitel/api/pkg/service/rest/users"
)

// Rest represents an implementation of a REST service
type Rest struct {
	Pods  svcdef.PodsService
	Users svcdef.UsersService
}

// New creates a new rest service instance
func New(cfg *config.Config) (*Rest, error) {
	podsService, err := podssvc.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed initializing pods REST service: %w", err)
	}

	usersService, err := userssvc.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed initializing users REST service: %w", err)
	}

	return &Rest{
		Pods:  podsService,
		Users: usersService,
	}, nil
}
