package cmd

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/generator"
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

	return generator.GenerateProject(ctx, &projectConfig)
}
