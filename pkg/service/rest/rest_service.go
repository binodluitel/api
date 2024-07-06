package rest

import (
	"fmt"

	"github.com/binodluitel/api/pkg/config"
	svcdef "github.com/binodluitel/api/pkg/service/definitions"
	streamsvc "github.com/binodluitel/api/pkg/service/rest/stream"
)

// Rest represents an implementation of a REST service
type Rest struct {
	Stream svcdef.StreamService
}

// New creates a new rest service instance
func New(cfg *config.Config) (*Rest, error) {
	streamService, err := streamsvc.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed initializing stream service: %w", err)
	}
	return &Rest{
		Stream: streamService,
	}, nil
}
