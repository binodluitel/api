package definitions

import (
	"context"
	"io"

	"github.com/binodluitel/api/pkg/models/pods"
)

//go:generate ../../../.build/bin/mockery --name=PodsService

// PodsService defines a methods related to k8s pods REST API service
type PodsService interface {
	GetLogs(context.Context, *pods.Logs) (io.ReadCloser, error)
}
