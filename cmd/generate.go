package cmd

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
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Run generator",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return runGenerateCmd(cmd.Context())
	},
}

func runGenerateCmd(ctx context.Context) error {
	var projectConfig config.ProjectConfig
	p := tea.NewProgram(tui.InitialModel(&projectConfig))
	if _, err := p.Run(); err != nil {
		return err
	}

	if projectConfig.Aborted {
		return nil
	}

	if err := validateProjectConfig(&projectConfig); err != nil {
		return err
	}

	infraSetConfig(&projectConfig)

	return generator.GenerateProject(ctx, &projectConfig)
}

func validateProjectConfig(projectConfig *config.ProjectConfig) error {
	if projectConfig.ProjectName == "" {
		return fmt.Errorf("project name is required")
	}
	if projectConfig.OpenAPIPath == "" {
		return fmt.Errorf("openapi path is required")
	}
	if projectConfig.ModuleName == "" {
		return fmt.Errorf("module name is required")
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

func infraSetConfig(cfg *config.ProjectConfig) {
	infra.GetDB(cfg.DB).Processor.SetConfig(cfg)

	for _, item := range infra.ListInfraInfos() {
		item.Processor.SetConfig(cfg)
	}
}
