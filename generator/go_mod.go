package generator

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/rom8726/airo/config"
)

//go:embed templates/go.mod.tmpl
var tmplGoMod string

func writeGoMod(cfg *config.ProjectConfig) error {
	tmpl, err := template.New("go_mod").Parse(tmplGoMod)
	if err != nil {
		return fmt.Errorf("parse template \"go_mod\" failed: %w", err)
	}

	goModFilePath := filepath.Join(cfg.ProjectName, "go.mod")
	fGoMod, err := os.Create(goModFilePath)
	if err != nil {
		return fmt.Errorf("create file \"%s\" failed: %w", goModFilePath, err)
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
