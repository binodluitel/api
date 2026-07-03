package rest

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	livenessctrl "github.com/binodluitel/api/pkg/api/rest/controllers/liveness"
	podsctrl "github.com/binodluitel/api/pkg/api/rest/controllers/pods"
	readinessctrl "github.com/binodluitel/api/pkg/api/rest/controllers/readiness"
	usersctrl "github.com/binodluitel/api/pkg/api/rest/controllers/users"
	"github.com/binodluitel/api/pkg/config"
	"github.com/binodluitel/api/pkg/log"
	restservice "github.com/binodluitel/api/pkg/service/rest"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
)

// Rest defines a REST application
type Rest struct {
	Engine     *gin.Engine
	readyCheck func(context.Context) error
}

func New(cfg *config.Config, rest *restservice.Rest) (*Rest, error) {
	gin.SetMode(cfg.API.Rest.Mode)
	engine := gin.New()
	engine.Use(
		gin.Recovery(), // recover REST API from any panics
		otelgin.Middleware(cfg.Application.Name),
	)

	// base router group
	router := engine.Group("")

	// add middlewares to base routes
	router.Use()

	// root path to return I AM A TEAPOT response code
	router.Any("/", func(c *gin.Context) { c.Status(http.StatusTeapot) })

	// v1 API router group
	v1Router := router.Group("v1")
	podsctrl.New(rest.Pods, v1Router)
	usersctrl.New(rest.Users, v1Router)

	// Health check endpoints for Kubernetes
	livenessctrl.New(router)
	readinessctrl.New(rest.Pods, router)
	var readyCheck func(context.Context) error
	if checker, ok := rest.Pods.(interface{ Ready(context.Context) error }); ok {
		readyCheck = checker.Ready
	} else {
		return nil, fmt.Errorf("pods service does not support readiness checks")
	}
	return &Rest{Engine: engine, readyCheck: readyCheck}, nil
}

// Run starts REST server and handles graceful shutdown
func (r *Rest) Run(ctx context.Context, cfg *config.Config, logger log.Logger, shutdownTimeout time.Duration) error {
	address := net.JoinHostPort(cfg.API.Rest.Host, cfg.API.Rest.Port)

	server := &http.Server{
		Addr:    address,
		Handler: r.Engine,
	}

	// Start server in a goroutine
	go func() {
		logger.Info("REST API server started", zap.String("address", address))
		var err error
		if cfg.API.Rest.TLS.Enable {
			err = server.ListenAndServeTLS(
				cfg.API.Rest.TLS.Server.CertPath,
				cfg.API.Rest.TLS.Server.KeyPath,
			)
		} else {
			err = server.ListenAndServe()
		}
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("REST API server error", zap.Error(err))
		}
	}()

	// Wait for context cancellation (signal or error)
	<-ctx.Done()
	logger.Info("REST API server shutdown initiated")

	// Graceful shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("REST API server shutdown error", zap.Error(err))
		return err
	}

	logger.Info("REST API server shut down gracefully")
	return nil
}
