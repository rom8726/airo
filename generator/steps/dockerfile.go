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

//go:embed templates/dockerfile.tmpl
var tmplDockerfile string

type DockerfileStep struct{}

func (DockerfileStep) Description() string {
	return "Create Dockerfile"
}

func (DockerfileStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
	tmpl, err := template.New("dockerfile").Parse(tmplDockerfile)
	if err != nil {
		return fmt.Errorf("parse template \"dockerfile\" failed: %w", err)
	}

	dockerfilePath := filepath.Join(projectDir(cfg), "Dockerfile")
	f, err := os.Create(dockerfilePath)
	if err != nil {
		return fmt.Errorf("create file %q failed: %w", dockerfilePath, err)
	}
	defer f.Close()

	type renderData struct {
		GoVersion string
	}
	data := renderData{
		GoVersion: goVersion,
	}

	if err := tmpl.Execute(f, data); err != nil {
		return fmt.Errorf("execute template \"dockerfile\" failed: %w", err)
	}

	return nil
}
