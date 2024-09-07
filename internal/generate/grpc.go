package generate

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/ryanreadbooks/prae/internal/config"
	"github.com/ryanreadbooks/prae/internal/pkg"
	"github.com/ryanreadbooks/prae/internal/pkg/log"
)

func renderInternalGrpc(cfg *config.Config) error {
	if err := doProtogen(cfg); err != nil {
		return fmt.Errorf("do protoc-gen: %w", err)
	}

	err := pkg.RenderTemplate("templates/internal/rpc/server.go",
		outAtInternalRpc(cfg.OutDir(), "server.go"),
		templates,
		cfg)

	return err
}

func doProtogen(cfg *config.Config) error {
	if len(cfg.Grpc.Inputs) == 0 {
		return fmt.Errorf("proto input directory empty")
	}

	files := make(map[string][]string)

	// traverse input to fetch all .proto files
	for _, root := range cfg.Grpc.Inputs {
		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !d.IsDir() && isProtoFile(path) {
				// simple check here
				if strings.HasPrefix(path, root) {
					files[root] = append(files[root], path)
				}
			}

			return nil
		})

		if err != nil {
			return fmt.Errorf("walk through %s: %w", root, err)
		}
	}

	// perform protoc-gen-xxx
	// protoc -I {path} xxxx --xx_out={z} --xx_opt=
	var builder strings.Builder
	builder.WriteString("protoc ")

	for _, extra := range cfg.Grpc.Extras {
		builder.WriteString(extra)
		builder.WriteByte(' ')
	}

	for dir, files := range files {
		builder.WriteString(fmt.Sprintf("-I %s ", dir))
		for _, file := range files {
			builder.WriteString(file)
			builder.WriteByte(' ')
		}
	}

	for _, plugin := range cfg.Grpc.Plugins {
		builder.WriteString(fmt.Sprintf("--%s_out=%s", plugin.Name, plugin.Out))
		if len(plugin.Out) != 0 {
			builder.WriteString(fmt.Sprintf(" --%s_opt=%s", plugin.Name, plugin.Opt))
		}
		builder.WriteByte(' ')
	}

	log.Info("%s\n", builder.String())
	cmds := strings.Split(strings.TrimSpace(builder.String()), " ")
	if len(cmds) <= 1 {
		return fmt.Errorf("internal error")
	}

	return pkg.ExecCommand(cmds[0], cmds[1:]...)
}

func isProtoFile(s string) bool {
	return filepath.Ext(s) == ".proto"
}

func pluginExecName(name string) string {
	return "protoc-gen-" + name
}
