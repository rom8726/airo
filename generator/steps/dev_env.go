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

//go:embed templates/makefile.tmpl
var tmplMakefile string

//go:embed templates/docker-compose.yml.tmpl
var tmplDockerComposeYml string

//go:embed templates/compose.env.example.tmpl
var tmplComposeEnv string

//go:embed templates/config.env.example.tmpl
var tmplConfigEnv string

//go:embed files/dev/dev.mk
var devMkFile []byte

type DevEnvStep struct{}

func (DevEnvStep) Description() string {
	return "Create Makefile"
}

func (DevEnvStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
	if err := doMakefile(cfg); err != nil {
		return err
	}

	if err := doEnvDir(cfg); err != nil {
		return err
	}

	return nil
}

func doMakefile(cfg *config.ProjectConfig) error {
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

func doEnvDir(cfg *config.ProjectConfig) error {
	dir := devDir(cfg)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}

	devMkFilePath := filepath.Join(dir, "dev.mk")
	if err := os.WriteFile(devMkFilePath, devMkFile, 0644); err != nil {
		return fmt.Errorf("failed to write dev.mk: %w", err)
	}

	if err := doDockerComposeYml(cfg); err != nil {
		return err
	}

	if err := doConfigEnvExample(cfg); err != nil {
		return err
	}

	if err := doComposeEnvExample(cfg); err != nil {
		return err
	}

	return nil
}

func doDockerComposeYml(cfg *config.ProjectConfig) error {
	tmpl, err := template.New("docker_compose").Parse(tmplDockerComposeYml)
	if err != nil {
		return fmt.Errorf("parse template \"docker_compose\" failed: %w", err)
	}

	dockerComposeYmlPath := filepath.Join(devDir(cfg), "docker-compose.yml")
	fDockerComposeYml, err := os.Create(dockerComposeYmlPath)
	if err != nil {
		return fmt.Errorf("create file %q failed: %w", dockerComposeYmlPath, err)
	}
	defer fDockerComposeYml.Close()

	type renderData struct {
		ProjectName string
		DB          infra.DBInfo
		Infras      []infra.InfraInfo
	}

	infras := make([]infra.InfraInfo, 0, len(cfg.UseInfra))
	for _, code := range cfg.UseInfra {
		item := infra.GetInfra(code)
		infras = append(infras, item)
	}

	data := renderData{
		ProjectName: cfg.ProjectName,
		DB:          infra.GetDB(cfg.DB),
		Infras:      infras,
	}

	if err := tmpl.Execute(fDockerComposeYml, data); err != nil {
		return fmt.Errorf("execute template \"docker_compose\" failed: %w", err)
	}

	return nil
}

func doConfigEnvExample(cfg *config.ProjectConfig) error {
	tmpl, err := template.New("config_env").Parse(tmplConfigEnv)
	if err != nil {
		return fmt.Errorf("parse template \"config_env\" failed: %w", err)
	}

	configEnvPath := filepath.Join(devDir(cfg), "config.env.example")
	fConfigEnv, err := os.Create(configEnvPath)
	if err != nil {
		return fmt.Errorf("create file %q failed: %w", configEnvPath, err)
	}
	defer fConfigEnv.Close()

	type renderData struct {
		DB     infra.DBInfo
		Infras []infra.InfraInfo
	}

	infras := make([]infra.InfraInfo, 0, len(cfg.UseInfra))
	for _, code := range cfg.UseInfra {
		item := infra.GetInfra(code)
		infras = append(infras, item)
	}

	data := renderData{
		DB:     infra.GetDB(cfg.DB),
		Infras: infras,
	}

	if err := tmpl.Execute(fConfigEnv, data); err != nil {
		return fmt.Errorf("execute template \"config_env\" failed: %w", err)
	}

	return nil
}

func doComposeEnvExample(cfg *config.ProjectConfig) error {
	tmpl, err := template.New("compose_env").Parse(tmplComposeEnv)
	if err != nil {
		return fmt.Errorf("parse template \"compose_env\" failed: %w", err)
	}

	composeEnvPath := filepath.Join(devDir(cfg), "compose.env.example")
	fComposeEnv, err := os.Create(composeEnvPath)
	if err != nil {
		return fmt.Errorf("create file %q failed: %w", composeEnvPath, err)
	}
	defer fComposeEnv.Close()

	type renderData struct {
		DB     infra.DBInfo
		Infras []infra.InfraInfo
	}

	infras := make([]infra.InfraInfo, 0, len(cfg.UseInfra))
	for _, code := range cfg.UseInfra {
		item := infra.GetInfra(code)
		infras = append(infras, item)
	}

	data := renderData{
		DB:     infra.GetDB(cfg.DB),
		Infras: infras,
	}

	if err := tmpl.Execute(fComposeEnv, data); err != nil {
		return fmt.Errorf("execute template \"compose_env\" failed: %w", err)
	}

	return nil
}
