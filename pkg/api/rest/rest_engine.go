package rest

import (
	"net"
	"net/http"

	podsctrl "github.com/binodluitel/api/pkg/api/rest/controllers/pods"
	usersctrl "github.com/binodluitel/api/pkg/api/rest/controllers/users"
	"github.com/binodluitel/api/pkg/config"
	restservice "github.com/binodluitel/api/pkg/service/rest"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// Rest defines a REST application
type Rest struct {
	Engine *gin.Engine
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
	return &Rest{engine}, nil
}

// Run starts a new REST listner
func (r *Rest) Run(cfg *config.Config) error {
	if cfg.API.Rest.TLS.Enable {
		return r.Engine.RunTLS(
			net.JoinHostPort(cfg.API.Rest.Host, cfg.API.Rest.Port),
			cfg.API.Rest.TLS.Server.CertPath,
			cfg.API.Rest.TLS.Server.KeyPath,
		)
	}
	return r.Engine.Run(net.JoinHostPort(cfg.API.Rest.Host, cfg.API.Rest.Port))
}
