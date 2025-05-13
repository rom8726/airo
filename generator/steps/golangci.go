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

//go:embed templates/golangci.yml.tmpl
var tmplGolangCI string

type GolangCIStep struct{}

func (GolangCIStep) Description() string {
	return "Create .golangci.yml"
}

func (GolangCIStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
	tmpl, err := template.New("golangci").Parse(tmplGolangCI)
	if err != nil {
		return fmt.Errorf("parse template \"golangci\" failed: %w", err)
	}

	golangciFilePath := filepath.Join(projectDir(cfg), ".golangci.yml")
	f, err := os.Create(golangciFilePath)
	if err != nil {
		return fmt.Errorf("create file %q failed: %w", golangciFilePath, err)
	}
	defer f.Close()

	type renderData struct {
		Module string
	}
	data := renderData{
		Module: cfg.ModuleName,
	}

	if err := tmpl.Execute(f, data); err != nil {
		return fmt.Errorf("execute template \"golangci\" failed: %w", err)
	}

	return nil
}
