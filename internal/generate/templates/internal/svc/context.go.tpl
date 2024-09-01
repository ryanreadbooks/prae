package svc

import (
	"{{ .Go.Module }}/internal/config"
	"{{ .Go.Module }}/internal/repo"
)

type ServiceContext struct {
	Config *config.Config
	Repo	 *repo.Repo
}

// Init a service object
func New(c *config.Config) *ServiceContext {
	r := repo.New(c)
	ctx := &ServiceContext{
		Config: c,
		Repo: 	r,
	}

	return ctx
}