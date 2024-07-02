package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

// prefix is the application config env var prefix
const prefix = ""

// Config contains configuration parameters
type Config struct {
	Application Application `envconfig:"app"`
	API         struct {
		Rest *Rest `envconfig:"rest"`
	} `envconfig:"api"`
	Log       Log       `envconfig:"log"`
	Telemetry Telemetry `envconfig:"telemetry"`
}

// Application configuration
type Application struct {
	Name      string `envconfig:"name" default:"api-service"`
	BuildTime string `envconfig:"build_time"`
	Version   string `envconfig:"version"`
	Git       struct {
		RefName string `envconfig:"ref_name"`
		RefSHA  string `envconfig:"ref_sha"`
	} `envconfig:"git"`
}

// Log is configuration information for application language
type Log struct {
	Level       string `envconfig:"level" default:"debug"`
	Development bool   `envconfig:"development" default:"true"`
	Encoding    string `envconfig:"encoding" default:"json"`
	TracerName  string `envconfig:"tracer_name" default:"unknown-log-tracer"`
}

// Rest is REST api configuration
type Rest struct {
	Host string `envconfig:"host" default:""`
	Mode string `envconfig:"mode" default:"debug"`
	Port string `envconfig:"port" default:"8080"`
	TLS  struct {
		Enable bool `envconfig:"enable" default:"false"`
		Server struct {
			CertFile string `envconfig:"cert_file"`
			KeyFile  string `envconfig:"key_file"`
		} `envconfig:"server"`
	} `envconfig:"tls"`
}

// Telemetry configurations
type Telemetry struct {
	Metrics Metrics `envconfig:"metrics"`
}

type Metrics struct {
	Host string `envconfig:"host" default:""`
	Port string `envconfig:"port"    default:"9090"`
	TLS  struct {
		Enable   bool   `envconfig:"enable" default:"false"`
		CertFile string `envconfig:"cert_file"`
		KeyFile  string `envconfig:"key_file"`
	} `envconfig:"tls"`
}

// MustGet will return config or panic
func MustGet() *Config {
	c, err := initialize()
	if err != nil {
		panic(err)
	}
	return c
}

// Get will return config or error
func Get() (*Config, error) {
	c, err := initialize()
	if err != nil {
		return nil, err
	}
	return c, nil
}

// global application configuration instance
var c *Config
var mutex sync.Mutex

// initialize application configuration
// this initializes the configuration only once if not already initialized
func initialize() (*Config, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if c != nil {
		return c, nil
	}
	c = new(Config)
	if err := envconfig.Process(prefix, c); err != nil {
		return nil, err
	}
	if err := c.validate(); err != nil {
		return nil, err
	}
	return c, nil
}

// validate information in the configuration
func (c *Config) validate() error {
	return nil
}
