package main

import (
	"context"
	"fmt"
	"net"

	restapi "github.com/binodluitel/api/pkg/api/rest"
	"github.com/binodluitel/api/pkg/config"
	"github.com/binodluitel/api/pkg/log"
	apimetrics "github.com/binodluitel/api/pkg/metrics"
	restservice "github.com/binodluitel/api/pkg/service/rest"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	group, ctx := errgroup.WithContext(ctx)
	// Initialize Application Configurations
	// ----------------------------------------------------------------------------------------
	cfg, err := config.Get()
	if err != nil {
		panic(fmt.Sprintf("failed initializing API application configuration, %s", err))
	}

	_, logger := log.Get(ctx)
	defer logger.Sync()
	logger.Info(" ----- Welcome to the API service example ----- ")
	logger.With([]zap.Field{
		zap.String("name", cfg.Application.Name),
		zap.String("version", cfg.Application.Version),
		zap.String("build_time", cfg.Application.BuildTime),
		zap.String("ref_name", cfg.Application.Git.RefName),
		zap.String("ref_sha", cfg.Application.Git.RefSHA),
	}...).Debug("Application build information")

	// Add build information to metrics
	apimetrics.BuildInfo.With(prometheus.Labels{
		"build_time":   cfg.Application.BuildTime,
		"version":      cfg.Application.Version,
		"git_ref_name": cfg.Application.Git.RefName,
		"git_ref_sha":  cfg.Application.Git.RefSHA,
	}).Set(1)
	if err := apimetrics.RegisterTo(prometheus.DefaultRegisterer); err != nil {
		logger.Panic(fmt.Sprintf("failed to register metrics, %s", err))
	}

	// Start Servers
	// ----------------------------------------------------------------------------------------
	// Start REST application
	group.Go(func() error {
		return startRestAPIService(ctx, cfg)
	})

	// Start metrics server
	group.Go(func() error {
		return startMetricsServer(ctx, cfg)
	})

	// Wait here
	if err := group.Wait(); err != nil {
		logger.Panic(fmt.Sprintf("failed starting API application service, %s", err))
	}
}

// startRestAPIService starts REST API service
func startRestAPIService(_ context.Context, cfg *config.Config) error {
	restService, err := restservice.New(cfg)
	if err != nil {
		return err
	}
	restEngine, err := restapi.New(cfg, restService)
	if err != nil {
		return err
	}
	if err := restEngine.Run(cfg); err != nil {
		return fmt.Errorf("failed to start REST API service, %s", err)
	}
	return nil
}

// startMetricsServer starts Prometheus metrics service server
func startMetricsServer(_ context.Context, cfg *config.Config) error {
	promHandler := promhttp.Handler()
	engine := gin.New()
	engine.GET("/metrics", func(c *gin.Context) {
		promHandler.ServeHTTP(c.Writer, c.Request)
	})
	metricsAddress := net.JoinHostPort(cfg.Telemetry.Metrics.Host, cfg.Telemetry.Metrics.Port)
	if cfg.Telemetry.Metrics.TLS.Enable {
		if err := engine.RunTLS(
			metricsAddress,
			cfg.Telemetry.Metrics.TLS.Server.CertPath,
			cfg.Telemetry.Metrics.TLS.Server.KeyPath,
		); err != nil {
			return fmt.Errorf("failed to start secured metrics server, %s", err)
		}
	}
	if err := engine.Run(metricsAddress); err != nil {
		return fmt.Errorf("failed to start metrics server, %s", err)
	}
	return nil
}
