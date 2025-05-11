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

//go:embed templates/main.go.tmpl
var tmplMainGo string

type MainGoStep struct{}

func (MainGoStep) Description() string {
	return "Create main.go"
}

func (MainGoStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
	tmpl, err := template.New("main_go").Parse(tmplMainGo)
	if err != nil {
		return fmt.Errorf("parse template \"main_go\" failed: %w", err)
	}

	mainGoFilePath := filepath.Join(projectDir(cfg), "main.go")
	fMainGo, err := os.Create(mainGoFilePath)
	if err != nil {
		return fmt.Errorf("create file \"%s\" failed: %w", mainGoFilePath, err)
	}
	defer fMainGo.Close()

	type renderData struct {
		ModuleName string
	}
	data := renderData{
		ModuleName: cfg.ModuleName,
	}

	if err := tmpl.Execute(fMainGo, data); err != nil {
		return fmt.Errorf("execute template \"main_go\" failed: %w", err)
	}

	return nil
}
