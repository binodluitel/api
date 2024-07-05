package stream

import (
	svcdef "github.com/binodluitel/api/pkg/service/definitions"
)

// Stream defines streaming service instance
type Stream struct{}

// New creates and returns a new user service instance
func New() svcdef.StreamService {
	return &Stream{}
}
