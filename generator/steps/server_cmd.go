package steps

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/generator/infra"
)

//go:embed templates/server_cmd.go.tmpl
var tmplServerCmd string

type ServerCmdStep struct {
	registry *infra.Registry
}

func NewServerCmdStep(registry *infra.Registry) *ServerCmdStep {
	return &ServerCmdStep{
		registry: registry,
	}
}

func (ServerCmdStep) Description() string {
	return "Create server command"
}

func (s ServerCmdStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
	dir := serverCmdDir(cfg)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}

	serverFilePath := filepath.Join(dir, "server.go")
	fServer, err := os.Create(serverFilePath)
	if err != nil {
		return fmt.Errorf("create file %q failed: %w", serverFilePath, err)
	}
	defer fServer.Close()

	tmpl, err := template.New("server_cmd").Parse(tmplServerCmd)
	if err != nil {
		return fmt.Errorf("parse template \"server_cmd\" failed: %w", err)
	}

	type RenderData struct {
		Module             string
		DBConfigFieldName  string
		HasSecurityHandler bool
	}
	data := RenderData{
		Module:             cfg.ModuleName,
		DBConfigFieldName:  s.registry.GetDB(cfg.DB).Processor.ConfigFieldName(),
		HasSecurityHandler: hasSecurityHandler(cfg),
	}

	if err := tmpl.Execute(fServer, data); err != nil {
		return fmt.Errorf("execute template \"server_cmd\" failed: %w", err)
	}

	return nil
}
