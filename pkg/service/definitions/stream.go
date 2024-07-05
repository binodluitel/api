package definitions

import (
	"context"

	"github.com/binodluitel/api/pkg/models"
)

//go:generate ../../../.build/bin/mockery --name=StreamService

// StreamService defines a methods for stream service
type StreamService interface {
	StreamLogs(context.Context, *models.StreamRequest) (*string, error)
}
