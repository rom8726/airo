package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/generator"
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
	//p := tea.NewProgram(tui.InitialModel(&projectConfig))
	//if _, err := p.Run(); err != nil {
	//	return err
	//}

	projectConfig.ProjectName = "proj1"
	projectConfig.ModuleName = "github.com/user/proj1"
	projectConfig.OpenAPIPath = "server.yml"
	projectConfig.UsePostgres = true
	projectConfig.UseRedis = true

	return generator.GenerateProject(ctx, &projectConfig)
}
