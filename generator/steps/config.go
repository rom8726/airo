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

//go:embed templates/config.go.tmpl
var tmplConfigGo string

type ConfigStep struct {
	reg *infra.Registry
}

func NewConfigStep(reg *infra.Registry) *ConfigStep {
	return &ConfigStep{
		reg: reg,
	}
}

func (ConfigStep) Description() string {
	return "Generate config package"
}

func (s ConfigStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
	dir := configDir(cfg)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}

	configFilePath := filepath.Join(dir, "config.go")
	fConfig, err := os.Create(configFilePath)
	if err != nil {
		return fmt.Errorf("create file %q failed: %w", configFilePath, err)
	}
	defer fConfig.Close()

	tmpl, err := template.New("config_go").Parse(tmplConfigGo)
	if err != nil {
		return fmt.Errorf("parse template \"config_go\" failed: %w", err)
	}

	infraInfos := make([]infra.InfraInfo, 0, len(cfg.UseInfra))
	for _, code := range cfg.UseInfra {
		infraInfos = append(infraInfos, s.reg.GetInfra(code))
	}

	type renderData struct {
		DB             infra.DBInfo
		Infras         []infra.InfraInfo
		UseRealtimeJWT bool
	}
	data := renderData{
		DB:             s.reg.GetDB(cfg.DB),
		Infras:         infraInfos,
		UseRealtimeJWT: cfg.UseRealtimeJWT,
	}

	if err := tmpl.Execute(fConfig, data); err != nil {
		return fmt.Errorf("execute template \"config_go\" failed: %w", err)
	}

	return nil
}
