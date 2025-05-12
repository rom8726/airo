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

//go:embed templates/makefile.tmpl
var tmplMakefile string

type MakefileStep struct{}

func (MakefileStep) Description() string {
	return "Create Makefile"
}

func (MakefileStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
	tmpl, err := template.New("makefile").Parse(tmplMakefile)
	if err != nil {
		return fmt.Errorf("parse template \"makefile\" failed: %w", err)
	}

	makefilePath := filepath.Join(projectDir(cfg), "Makefile")
	fMakefile, err := os.Create(makefilePath)
	if err != nil {
		return fmt.Errorf("create file %q failed: %w", makefilePath, err)
	}
	defer fMakefile.Close()

	type renderData struct {
		Module      string
		ProjectName string
	}

	data := renderData{
		Module:      cfg.ModuleName,
		ProjectName: cfg.ProjectName,
	}

	if err := tmpl.Execute(fMakefile, data); err != nil {
		return fmt.Errorf("execute template \"makefile\" failed: %w", err)
	}

	return nil
}
