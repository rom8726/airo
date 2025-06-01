package steps

import (
	"context"
	_ "embed"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"text/template"

	"github.com/rom8726/airo/config"
)

//go:embed files/tests/runner/env.go.example
var testsRunnerEnvGo []byte

//go:embed files/tests/runner/migrate_postgresql.go.example
var testsRunnerMigratePostgresGo []byte

//go:embed files/tests/runner/migrate_mysql.go.example
var testsRunnerMigrateMySQLGo []byte

//go:embed files/tests/runner/runner.go.tmpl
var testsRunnerGoTemplate string

type TestyStep struct{}

func (TestyStep) Description() string {
	return "Generate Testy files"
}

func (TestyStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
	if !cfg.UseTesty {
		slog.Info("skipped")

		return nil
	}

	dir := testsRunnerDir(cfg)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}

	envFilePath := filepath.Join(dir, "env.go")
	if err := os.WriteFile(envFilePath, testsRunnerEnvGo, 0644); err != nil {
		return fmt.Errorf("write file %q failed: %w", envFilePath, err)
	}

	migrateFilePath := filepath.Join(dir, "migrate.go")
	var migrateFileContent []byte
	switch cfg.DB {
	case config.DBTypeMySQL:
		migrateFileContent = testsRunnerMigrateMySQLGo
	case config.DBTypePostgres:
		migrateFileContent = testsRunnerMigratePostgresGo
	default:
		// Default to PostgreSQL for backward compatibility
		migrateFileContent = testsRunnerMigratePostgresGo
	}

	if err := os.WriteFile(migrateFilePath, migrateFileContent, 0644); err != nil {
		return fmt.Errorf("write file %q failed: %w", migrateFilePath, err)
	}

	runnerFilePath := filepath.Join(dir, "runner.go")
	fRunner, err := os.Create(runnerFilePath)
	if err != nil {
		return fmt.Errorf("create file %q failed: %w", runnerFilePath, err)
	}
	defer fRunner.Close()

	type renderData struct {
		Module   string
		UseInfra []string
		DB       string
	}
	data := renderData{
		Module:   cfg.ModuleName,
		UseInfra: cfg.UseInfra,
		DB:       cfg.DB,
	}

	tmpl, err := template.New("test_runner").Parse(testsRunnerGoTemplate)
	if err != nil {
		return fmt.Errorf("parse template \"test_runner\" failed: %w", err)
	}

	if err := tmpl.Execute(fRunner, data); err != nil {
		return fmt.Errorf("execute template \"test_runner\" failed: %w", err)
	}

	return nil
}
