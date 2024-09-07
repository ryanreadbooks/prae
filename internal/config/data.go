package config

import (
	"fmt"
	"os"

	"github.com/ryanreadbooks/prae/internal/pkg/log"
	"github.com/ryanreadbooks/prae/internal/ver"
)

const (
	StyleZero = "zero"
)

// configuration file representation
type Config struct {
	Version string `yaml:"version"` // prae version
	AppName string `yaml:"app"`     // project name
	Go      *Go    `yaml:"go"`      // go configuration
	Style   string `yaml:"style"`   // which go framework to use
	Grpc    *Grpc  `yaml:"grpc"`

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

	return nil
}

func (c *Config) ZeroStyle() bool {
	return c.Style == StyleZero
}

func (c *Config) OutDir() string {
	return c.outdir
}

func (c *Config) HasGrpc() bool {
	return c.Grpc != nil
}

func (c *Config) HasHttp() bool {
	return false
}

type Go struct {
	Version string `yaml:"version"`
	Module  string `yaml:"module"`
	Tidy    bool   `yaml:"tidy"`
}

type Grpc struct {
	Inputs  []string     `yaml:"inputs"` // input directory
	Plugins []GrpcPlugin `yaml:"plugins"`
	Extras  []string     `yaml:"extras"`
}

func (c *Grpc) Check() error {
	if len(c.Inputs) == 0 {
		return fmt.Errorf("proto inputs should not be empty")
	}

	for _, p := range c.Plugins {
		if err := p.Check(); err != nil {
			return err
		}
	}
	return nil
}

type GrpcPlugin struct {
	Name string `yaml:"name"`
	Out  string `yaml:"out"`
	Opt  string `yaml:"opt"`
}

func (p *GrpcPlugin) Check() error {
	if p.Name == "" {
		return fmt.Errorf("plugin name should not be empty")
	}

	if p.Out == "" {
		return fmt.Errorf("plugin out should not be empty")
	}

	return nil
}
