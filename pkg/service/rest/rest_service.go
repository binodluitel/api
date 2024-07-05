package rest

import (
	svcdef "github.com/binodluitel/api/pkg/service/definitions"
	streamsvc "github.com/binodluitel/api/pkg/service/rest/stream"
)

// Rest represents an implementation of a REST service
type Rest struct {
	Stream svcdef.StreamService
}

// New creates a new rest service instance
func New() (*Rest, error) {
	return &Rest{
		Stream: streamsvc.New(),
	}, nil
}
