package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/generator"
	"github.com/rom8726/airo/generator/infra"
	"github.com/rom8726/airo/tui"
	"github.com/rom8726/airo/validate"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Run generator",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return runGenerateCmd(cmd.Context())
	},
}

func runGenerateCmd(ctx context.Context) error {
	registry := infra.NewRegistry(
		// DBs
		infra.WithPostgres(),
		infra.WithMySQL(),
		infra.WithMongo(),
		// Infra
		infra.WithRedis(),
		infra.WithKafka(),
		infra.WithElasticsearch(),
		infra.WithEtcd(),
		infra.WithNats(),
		infra.WithMemcache(),
		infra.WithRabbitMQ(),
		infra.WithAerospike(),
	)

	var projectConfig config.ProjectConfig
	p := tea.NewProgram(tui.InitialModel(&projectConfig, registry))
	if _, err := p.Run(); err != nil {
		return err
	}

	if projectConfig.Aborted {
		return nil
	}

	if err := validateProjectConfig(&projectConfig); err != nil {
		return err
	}

	registry.UpdateConfig(&projectConfig)

	gen := generator.New(registry)

	return gen.GenerateProject(ctx, &projectConfig)
}

func validateProjectConfig(projectConfig *config.ProjectConfig) error {
	if err := validate.ValidateProjectName(projectConfig.ProjectName); err != nil {
		return fmt.Errorf("project name is invalid: %w", err)
	}
	if err := validate.ValidateModuleName(projectConfig.ModuleName); err != nil {
		return fmt.Errorf("module name is invalid: %w", err)
	}
	if projectConfig.OpenAPIPath == "" {
		return fmt.Errorf("openapi path is required")
	}
	if projectConfig.DB == "" {
		return fmt.Errorf("db type is required")
	}

	dir := filepath.Join(".", projectConfig.ProjectName)
	if _, err := os.Stat(dir); err == nil {
		return fmt.Errorf("project directory %q already exists", dir)
	}

	if _, err := os.Stat(projectConfig.OpenAPIPath); os.IsNotExist(err) {
		return err
	}

	return nil
}
