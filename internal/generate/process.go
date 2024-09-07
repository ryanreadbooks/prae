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
var templates embed.FS

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

	if cfg.Go.Tidy {
		log.Info("performing go mod tidy\n")
		pkg.GoModTidy()
	}

	return nil
}

func outpath(outdir string, name ...string) string {
	e := make([]string, 0, len(name)+1)
	e = append(e, outdir)
	e = append(e, name...)
	return filepath.Join(e...)
}

func outAtEtc(outdir string, name ...string) string {
	e := make([]string, 0, len(name)+1)
	e = append(e, outdir)
	e = append(e, "etc")
	e = append(e, name...)
	return filepath.Join(e...)
}

func outAtInternal(outdir string, name ...string) string {
	e := make([]string, 0, len(name)+1)
	e = append(e, outdir)
	e = append(e, "internal")
	e = append(e, name...)
	return filepath.Join(e...)
}

func outAtInternalRepo(outdir string, name ...string) string {
	e := make([]string, 0, len(name)+2)
	e = append(e, outdir)
	e = append(e, "internal", "repo")
	e = append(e, name...)
	return filepath.Join(e...)
}

func outAtInternalConfig(outdir string, name ...string) string {
	e := make([]string, 0, len(name)+2)
	e = append(e, outdir)
	e = append(e, "internal", "config")
	e = append(e, name...)
	return filepath.Join(e...)
}

func outAtInternalSvc(outdir string, name ...string) string {
	e := make([]string, 0, len(name)+2)
	e = append(e, outdir)
	e = append(e, "internal", "svc")
	e = append(e, name...)
	return filepath.Join(e...)
}

func outAtInternalHttp(outdir string, name ...string) string {
	e := make([]string, 0, len(name)+2)
	e = append(e, outdir)
	e = append(e, "internal", "http")
	e = append(e, name...)
	return filepath.Join(e...)
}

func outAtInternalRpc(outdir string, name ...string) string {
	e := make([]string, 0, len(name)+2)
	e = append(e, outdir)
	e = append(e, "internal", "rpc")
	e = append(e, name...)
	return filepath.Join(e...)
}

func initProject(cfg *config.Config) error {
	log.Info("init project in %s\n", cfg.AppName)

	// go mod init
	err := pkg.RenderTemplate("templates/go.mod", outpath(cfg.OutDir(), "go.mod"), templates, cfg)
	if err != nil {
		return fmt.Errorf("go mod init: %w", err)
	}

	return nil
}

func renderCmd(cfg *config.Config) error {
	err := pkg.RenderTemplate("templates/cmd/main.go", outpath(cfg.OutDir(), "cmd", "main.go"),
		templates,
		cfg)
	if err != nil {
		return err
	}

	return nil
}

func renderEtc(cfg *config.Config) error {
	err := pkg.RenderTemplate("templates/etc/cfg.yml",
		outAtEtc(cfg.OutDir(), cfg.AppName+".yaml"),
		templates,
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
		outAtInternalConfig(cfg.OutDir(), "config.go"),
		templates,
		cfg)

	return err
}

func renderInternalRepo(cfg *config.Config) error {
	err := pkg.RenderTemplate("templates/internal/repo/repo.go",
		outAtInternalRepo(cfg.OutDir(), "repo.go"),
		templates,
		cfg)

	return err
}

func renderInternalSvc(cfg *config.Config) error {
	err := pkg.RenderTemplate("templates/internal/svc/context.go",
		outAtInternalSvc(cfg.OutDir(), "context.go"),
		templates,
		cfg)

	return err
}

func renderInternalHttpGrpc(cfg *config.Config) error {
	if cfg.HasHttp() {
		err := renderInternalHttp(cfg)
		if err != nil {
			return fmt.Errorf("render http: %w", err)
		}
	}
	if cfg.HasGrpc() {
		err := renderInternalGrpc(cfg)
		if err != nil {
			return fmt.Errorf("render grpc: %w", err)
		}
	}

	return nil
}

func renderInternalHttp(cfg *config.Config) error {
	err := pkg.RenderTemplate("templates/internal/http/server.go",
		outAtInternalHttp(cfg.OutDir(), "server.go"),
		templates,
		cfg)

	return err
}
