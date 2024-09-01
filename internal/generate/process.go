package generate

import (
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/ryanreadbooks/prae/internal/config"
	"github.com/ryanreadbooks/prae/internal/pkg"
	"github.com/ryanreadbooks/prae/internal/pkg/log"
)

//go:embed templates
var fs embed.FS

func genFile() error {

	return nil
}

func Handle(cfg *config.Config) error {
	if len(cfg.AppName) == 0 {
		return fmt.Errorf("app name is empty")
	}

	c, _ := json.Marshal(cfg)
	log.Info("%s\n", c)

	if err := initProject(cfg); err != nil {
		return err
	}

	if err := renderInternal(cfg); err != nil {
		return err
	}

	if err := renderCmd(cfg); err != nil {
		return err
	}

	if err := renderEtc(cfg); err != nil {
		return err
	}

	return nil
}

func initProject(cfg *config.Config) error {
	log.Info("init project in %s\n", cfg.AppName)
	err := pkg.EnsureDir(cfg.AppName)
	if err != nil {
		return fmt.Errorf("create project dir: %w", err)
	}

	// go mod init
	err = pkg.RenderTemplate("templates/go.mod", filepath.Join(cfg.AppName, "go.mod"), fs, cfg)
	if err != nil {
		return fmt.Errorf("go mod init: %w", err)
	}

	return nil
}

func renderCmd(cfg *config.Config) error {
	err := pkg.RenderTemplate("templates/cmd/main.go", filepath.Join(cfg.AppName, "cmd/main.go"),
		fs,
		cfg)
	if err != nil {
		return err
	}

	return nil
}

func renderEtc(cfg *config.Config) error {
	err := pkg.RenderTemplate("templates/etc/cfg.yml",
		filepath.Join(cfg.AppName, "etc", cfg.AppName+".yaml"),
		fs,
		cfg)

	return err
}

func renderInternal(cfg *config.Config) error {
	if err := renderInternalConfig(cfg); err != nil {
		return err
	}

	if err := renderInternalRepo(cfg); err != nil {
		return err
	}

	if err := renderInternalSvc(cfg); err != nil {
		return err
	}

	if err := renderInternalHttpGrpc(cfg); err != nil {
		return err
	}

	return nil
}

func renderInternalConfig(cfg *config.Config) error {
	err := pkg.RenderTemplate("templates/internal/config/config.go",
		filepath.Join(cfg.AppName, "internal/config/config.go"),
		fs,
		cfg)

	return err
}

func renderInternalRepo(cfg *config.Config) error {
	err := pkg.RenderTemplate("templates/internal/repo/repo.go",
		filepath.Join(cfg.AppName, "internal/repo/repo.go"), fs,
		cfg)

	return err
}

func renderInternalSvc(cfg *config.Config) error {
	err := pkg.RenderTemplate("templates/internal/svc/context.go",
		filepath.Join(cfg.AppName, "internal/svc/context.go"), fs,
		cfg)

	return err
}

func renderInternalHttpGrpc(cfg *config.Config) error {
	if cfg.ServiceTypeHttp() {
		return renderInternalHttp(cfg)
	} else if cfg.ServiceTypeGrpc() {
		return renderInternalGrpc(cfg)
	} else {
		if err := renderInternalHttp(cfg); err != nil {
			return fmt.Errorf("gen http: %w", err)
		}
		if err := renderInternalGrpc(cfg); err != nil {
			return fmt.Errorf("gen grpc: %w", err)
		}
	}

	return nil
}

func renderInternalHttp(cfg *config.Config) error {
	err := pkg.RenderTemplate("templates/internal/http/server.go",
		filepath.Join(cfg.AppName, "internal/http/server.go"),
		fs,
		cfg)

	return err
}

func renderInternalGrpc(cfg *config.Config) error {
	err := pkg.RenderTemplate("templates/internal/rpc/server.go",
		filepath.Join(cfg.AppName, "internal/rpc/server.go"),
		fs,
		cfg)

	return err
}
