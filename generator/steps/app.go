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

const (
	securityHandlerFileName = "oas_security_gen.go"
)

//go:embed templates/app.go.tmpl
var tmplAppGo string

type AppStep struct {
	reg *infra.Registry
}

func NewAppStep(reg *infra.Registry) *AppStep {
	return &AppStep{
		reg: reg,
	}
}

func (AppStep) Description() string {
	return "Create app.go"
}

func (s AppStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
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

	infraInfos := make([]infra.InfraInfo, 0, len(cfg.UseInfra))
	for _, code := range cfg.UseInfra {
		infraInfos = append(infraInfos, s.reg.GetInfra(code))
	}

	type renderData struct {
		Module             string
		DB                 infra.DBInfo
		Infras             []infra.InfraInfo
		HasSecurityHandler bool
	}
	data := renderData{
		Module:             cfg.ModuleName,
		DB:                 s.reg.GetDB(cfg.DB),
		Infras:             infraInfos,
		HasSecurityHandler: hasSecurityHandler(cfg),
	}

	if err := tmpl.Execute(fApp, data); err != nil {
		return fmt.Errorf("execute template \"app_go\" failed: %w", err)
	}

	return nil
}
