package metrics

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

// Heartbeat is application heartbeat that
// continuously beats at configurable fixed interval.
// It emits few metrics every configured intervals to say
// that the app is alive and provide the stats
func Heartbeat(ctx context.Context, meter metric.Meter) (context.Context, error) {
	// TODO: Implement the heartbeat function
	return ctx, nil
}
