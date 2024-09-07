package config

import (
	"embed"
	_ "embed"
	"fmt"

	"github.com/ryanreadbooks/prae/internal/pkg"
	"github.com/ryanreadbooks/prae/internal/ver"
)

const (
	outpath = "prae.yml"
)

//go:embed templates
var fs embed.FS

func Handle(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("prae project name should not be empty")
	}

	var data = Config{
		Version:     ver.Major(),
		AppName:     name,
		Go: &Go{
			Version: pkg.CollectGoVersion(),
			Module:  name,
		},
		Style: StyleZero,
	}

	return pkg.RenderTemplate("templates/config.yml", data.AppName+".yml", fs, data)
}
