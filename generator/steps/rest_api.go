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

//go:embed templates/rest_api.go.tmpl
var tmplApiGo string

type RestAPIStep struct{}

func (RestAPIStep) Description() string {
	return "Create REST API service component"
}

func (RestAPIStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
	dir := restAPIDir(cfg)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}

	apiFilePath := filepath.Join(dir, "api.go")
	fApi, err := os.Create(apiFilePath)
	if err != nil {
		return fmt.Errorf("create file %q failed: %w", apiFilePath, err)
	}
	defer fApi.Close()

	tmpl, err := template.New("api_go").Parse(tmplApiGo)
	if err != nil {
		return fmt.Errorf("parse template \"api_go\" failed: %w", err)
	}

	type renderData struct {
		Module string
	}
	data := renderData{
		Module: cfg.ModuleName,
	}

	if err := tmpl.Execute(fApi, data); err != nil {
		return fmt.Errorf("execute template \"api_go\" failed: %w", err)
	}

	return nil
}
