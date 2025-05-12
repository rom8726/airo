package steps

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/rom8726/airo/config"
)

const (
	securityHandlerFileName = "oas_security_gen.go"
)

//go:embed templates/app.go.tmpl
var tmplAppGo string

type AppStep struct{}

func (AppStep) Description() string {
	return "Create app.go"
}

func (AppStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
	dir := internalDir(cfg)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}

	appFilePath := filepath.Join(dir, "app.go")
	fApp, err := os.Create(appFilePath)
	if err != nil {
		return fmt.Errorf("create file %q failed: %w", appFilePath, err)
	}
	defer fApp.Close()

	tmpl, err := template.New("app_go").Parse(tmplAppGo)
	if err != nil {
		return fmt.Errorf("parse template \"app_go\" failed: %w", err)
	}

	type renderData struct {
		Module             string
		UsePostgres        bool
		UseRedis           bool
		HasSecurityHandler bool
	}
	data := renderData{
		Module:             cfg.ModuleName,
		UsePostgres:        cfg.DB == config.DBTypePostgres,
		UseRedis:           cfg.UseRedis,
		HasSecurityHandler: hasSecurityHandler(cfg),
	}

	if err := tmpl.Execute(fApp, data); err != nil {
		return fmt.Errorf("execute template \"app_go\" failed: %w", err)
	}

	return nil
}
