package rest

import (
	"net"
	"net/http"

	"github.com/binodluitel/api/pkg/config"
	restservice "github.com/binodluitel/api/pkg/service/rest"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// Rest defines a REST application
type Rest struct {
	Engine *gin.Engine
}

func New(restService *restservice.Rest) (*Rest, error) {
	cfg := config.MustGet()
	gin.SetMode(cfg.API.Rest.Mode)
	engine := gin.New()
	engine.Use(
		gin.Recovery(), // recover REST API from any panics
		otelgin.Middleware(cfg.Application.Name),
	)

	// base router group
	routes := engine.Group("")

	// add middlewares to base routes
	routes.Use()

	// root path to return I AM A TEAPOT response code
	routes.Any("/", func(c *gin.Context) { c.Status(http.StatusTeapot) })

	return &Rest{engine}, nil
}

// Run starts a new REST listner
func (r *Rest) Run() error {
	cfg := config.MustGet()
	if cfg.API.Rest.TLS.Enable {
		return r.Engine.RunTLS(
			net.JoinHostPort(cfg.API.Rest.Host, cfg.API.Rest.Port),
			cfg.API.Rest.TLS.Server.CertFile,
			cfg.API.Rest.TLS.Server.KeyFile,
		)
	}
	return r.Engine.Run(net.JoinHostPort(cfg.API.Rest.Host, cfg.API.Rest.Port))
}
