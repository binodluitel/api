package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
		_, _ = fmt.Fprintf(os.Stderr, "failed initializing API application configuration: %v\n", err)
		os.Exit(1)
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
		logger.Error("failed to register metrics", zap.Error(err))
		os.Exit(1)
	}

	// Setup signal handling for graceful shutdown
	// ----------------------------------------------------------------------------------------
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	// Convert config value (seconds) to duration
	shutdownTimeout := time.Duration(cfg.Application.GracefulShutdownTimeout) * time.Second

	// Start Servers
	// ----------------------------------------------------------------------------------------
	// Start REST application
	group.Go(func() error {
		return startRestAPIService(ctx, cfg, logger, shutdownTimeout)
	})

	// Start metrics server
	group.Go(func() error {
		return startMetricsServer(ctx, cfg, logger, shutdownTimeout)
	})

	// Handle graceful shutdown
	group.Go(func() error {
		sig := <-sigChan
		logger.Info("shutdown signal received", zap.String("signal", sig.String()))
		return nil
	})

	// Wait for all goroutines to complete
	if err := group.Wait(); err != nil {
		logger.Error("application error", zap.Error(err))
		os.Exit(1)
	}

	logger.Info("application shut down gracefully")
}

// startRestAPIService starts REST API service with graceful shutdown
func startRestAPIService(ctx context.Context, cfg *config.Config, logger log.Logger, shutdownTimeout time.Duration) error {
	restService, err := restservice.New(cfg)
	if err != nil {
		return fmt.Errorf("failed to create REST service: %w", err)
	}
	restEngine, err := restapi.New(cfg, restService)
	if err != nil {
		return fmt.Errorf("failed to create REST engine: %w", err)
	}
	if err := restEngine.Run(ctx, cfg, logger, shutdownTimeout); err != nil {
		return fmt.Errorf("failed to start REST API service: %w", err)
	}
	return nil
}

// startMetricsServer starts Prometheus metrics service server with graceful shutdown
func startMetricsServer(ctx context.Context, cfg *config.Config, logger log.Logger, shutdownTimeout time.Duration) error {
	promHandler := promhttp.Handler()
	engine := gin.New()
	engine.GET("/metrics", func(c *gin.Context) {
		promHandler.ServeHTTP(c.Writer, c.Request)
	})
	metricsAddress := net.JoinHostPort(cfg.Telemetry.Metrics.Host, cfg.Telemetry.Metrics.Port)

	var server *http.Server
	if cfg.Telemetry.Metrics.TLS.Enable {
		server = &http.Server{
			Addr:    metricsAddress,
			Handler: engine,
		}
		go func() {
			if err := server.ListenAndServeTLS(
				cfg.Telemetry.Metrics.TLS.Server.CertPath,
				cfg.Telemetry.Metrics.TLS.Server.KeyPath,
			); err != nil && !errors.Is(err, http.ErrServerClosed) {
				logger.Error("metrics server error", zap.Error(err))
			}
		}()
	} else {
		server = &http.Server{
			Addr:    metricsAddress,
			Handler: engine,
		}
		go func() {
			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				logger.Error("metrics server error", zap.Error(err))
			}
		}()
	}

	logger.Info("metrics server started", zap.String("address", metricsAddress))

	<-ctx.Done()
	logger.Info("metrics server shutdown initiated")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("metrics server shutdown error", zap.Error(err))
		return err
	}
	logger.Info("metrics server shut down gracefully")
	return nil
}
