package config

import (
	"fmt"
	"os"

	"github.com/ryanreadbooks/prae/internal/pkg"
	"github.com/ryanreadbooks/prae/internal/pkg/log"
	"github.com/ryanreadbooks/prae/internal/ver"
)

const (
	StyleZero = "zero"
)

const (
	ServiceTypeGrpc  = "grpc"
	ServiceTypeHttp  = "http"
	ServiceTypeBoth  = "grpc+http"
	ServiceTypeBoth2 = "http+grpc"
)

var (
	validStyles = map[string]struct{}{
		StyleZero: {},
	}

	validServiceTypes = map[string]struct{}{
		ServiceTypeGrpc:  {},
		ServiceTypeHttp:  {},
		ServiceTypeBoth:  {},
		ServiceTypeBoth2: {},
	}
)

// configuration file representation
type Config struct {
	Version     string `yaml:"version"` // prae version
	AppName     string `yaml:"app"`     // project name
	ServiceType string `yaml:"type"`    // service type
	Go          *Go    `yaml:"go"`      // go configuration
	Style       string `yaml:"style"`   // which go framework to use

	// internal use
	outdir string `yaml:"-"`
}

func NewConfig() (*Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return &Config{outdir: wd}, nil
}

// Check config validatity
func (c *Config) Check() error {
	// check version
	if c.Version != ver.Major() {
		log.Warn("version not match, %s != %s", c.Version, ver.Major())
	}

	// check style
	if err := validStyle(c.Style); err != nil {
		return err
	}

	if err := validServiceType(c.ServiceType); err != nil {
		return err
	}

	return nil
}

func validStyle(s string) error {
	if _, ok := validStyles[s]; ok {
		return nil
	}

	return fmt.Errorf("style not supported, see: %v", pkg.MapKeys(validStyles))
}

func (c *Config) ZeroStyle() bool {
	return c.Style == StyleZero
}

func validServiceType(s string) error {
	if _, ok := validServiceTypes[s]; ok {
		return nil
	}

	return fmt.Errorf("type not supported, see: %v", pkg.MapKeys(validServiceTypes))
}

func (c *Config) ServiceTypeGrpc() bool {
	return c.ServiceType == ServiceTypeGrpc
}

func (c *Config) ServiceTypeHttp() bool {
	return c.ServiceType == ServiceTypeHttp
}

func (c *Config) ServiceTypeBoth() bool {
	return c.ServiceType == ServiceTypeBoth || c.ServiceType == ServiceTypeBoth2
}

func (c *Config) ServiceTypeHasHttp() bool {
	return c.ServiceTypeHttp() || c.ServiceTypeBoth()
}

func (c *Config) ServiceTypeHasGrpc() bool {
	return c.ServiceTypeGrpc() || c.ServiceTypeBoth()
}

func (c *Config) OutDir() string {
	return c.outdir
}

type Go struct {
	Version string `yaml:"version"`
	Module  string `yaml:"module"`
	Tidy    bool   `yaml:"tidy"`
}
