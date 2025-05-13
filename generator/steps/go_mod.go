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

//go:embed templates/go.mod.tmpl
var tmplGoMod string

type GoModStep struct{}

func (GoModStep) Description() string {
	return "Create go.mod"
}

func (GoModStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
	tmpl, err := template.New("go_mod").Parse(tmplGoMod)
	if err != nil {
		return fmt.Errorf("parse template \"go_mod\" failed: %w", err)
	}

	goModFilePath := filepath.Join(projectDir(cfg), "go.mod")
	fGoMod, err := os.Create(goModFilePath)
	if err != nil {
		return fmt.Errorf("create file %q failed: %w", goModFilePath, err)
	}
	defer fGoMod.Close()

	type renderData struct {
		ModuleName string
		GoVersion  string
	}
	data := renderData{
		ModuleName: cfg.ModuleName,
		GoVersion:  goVersion,
	}

	if err := tmpl.Execute(fGoMod, data); err != nil {
		return fmt.Errorf("execute template \"go_mod\" failed: %w", err)
	}

	return nil
}
