package main

import (
	"context"
	"fmt"
	"net"

	restapi "github.com/binodluitel/api/pkg/api/rest"
	"github.com/binodluitel/api/pkg/config"
	"github.com/binodluitel/api/pkg/log"
	restservice "github.com/binodluitel/api/pkg/service/rest"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	group, ctx := errgroup.WithContext(ctx)
	_, logger := log.Get(ctx)
	defer logger.Sync()
	logger.Info(" ----- Welcome to the API service example ----- ")

	// Initialize Application Configurations
	// ----------------------------------------------------------------------------------------
	_, err := config.Get()
	if err != nil {
		logger.Panic(fmt.Sprintf("failed initializing API application configuration, %s", err))
	}

	// Start Servers
	// ----------------------------------------------------------------------------------------
	// Start REST application
	group.Go(func() error {
		return startRestAPIService(ctx)
	})

	// Start metrics server
	group.Go(func() error {
		return startMetricsServer(ctx)
	})

	// Wait here
	if err := group.Wait(); err != nil {
		logger.Panic(fmt.Sprintf("failed starting API application service, %s", err))
	}
}

// startRestAPIService starts REST API service
func startRestAPIService(_ context.Context) error {
	restService, err := restservice.New()
	if err != nil {
		return err
	}
	restEngine, err := restapi.New(restService)
	if err != nil {
		return err
	}
	if err := restEngine.Run(); err != nil {
		return fmt.Errorf("failed to start REST API service, %s", err)
	}
	return nil
}

// startMetricsServer starts Prometheus metrics service server
func startMetricsServer(_ context.Context) error {
	cfg := config.MustGet()
	promHandler := promhttp.Handler()
	engine := gin.New()
	engine.GET("/metrics", func(c *gin.Context) {
		promHandler.ServeHTTP(c.Writer, c.Request)
	})
	metricsAddress := net.JoinHostPort(cfg.Telemetry.Metrics.Host, cfg.Telemetry.Metrics.Port)
	if cfg.Telemetry.Metrics.TLS.Enable {
		if err := engine.RunTLS(
			metricsAddress,
			cfg.Telemetry.Metrics.TLS.CertFile,
			cfg.Telemetry.Metrics.TLS.KeyFile,
		); err != nil {
			return fmt.Errorf("failed to start secured metrics server, %s", err)
		}
	}
	if err := engine.Run(metricsAddress); err != nil {
		return fmt.Errorf("failed to start metrics server, %s", err)
	}
	return nil
}
