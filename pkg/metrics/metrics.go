package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "api_example"
)

// Metrics used throughout the application for troubleshooting and health indicators.
var (
	BuildInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "build_info",
			Help:      "Build Information",
		},
		[]string{
			"build_time",
			"version",
			"git_ref_name",
			"git_ref_sha",
		},
	)
)

// RegisterTo registers metrics to a desired target Registerer.
func RegisterTo(target prometheus.Registerer) error {
	collectors := []prometheus.Collector{
		BuildInfo,
	}
	for _, c := range collectors {
		if err := target.Register(c); err != nil {
			return fmt.Errorf("error registering %s metrics; %w", namespace, err)
		}
	}
	return nil
}
