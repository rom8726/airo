package generator

import (
	"context"
	"fmt"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/generator/infra"
	"github.com/rom8726/airo/generator/steps"
)

type Generator struct {
	registry *infra.Registry
}

func New(registry *infra.Registry) *Generator {
	return &Generator{
		registry: registry,
	}
}

func (g *Generator) GenerateProject(ctx context.Context, cfg *config.ProjectConfig) error {
	steps := []Step{
		steps.RootDirStep{},
		steps.GoModStep{},
		steps.SpecsStep{},
		steps.OGenStep{},
		steps.PkgStep{},
		steps.NewConfigStep(g.registry),
		steps.RestAPIStep{},
		steps.NewAppStep(g.registry),
		steps.NewServerCmdStep(g.registry),
		steps.MainGoStep{},
		steps.NewMigrateStep(g.registry),
		steps.DockerfileStep{},
		steps.NewDevEnvStep(g.registry),
		steps.GolangCIStep{},
		steps.GoModTidyStep{},
	}

	for _, step := range steps {
		if err := ctx.Err(); err != nil {
			return err
		}

		fmt.Printf("%s... ", step.Description())
		if err := step.Do(ctx, cfg); err != nil {
			return err
		}
		fmt.Println("done")
	}

	return nil
}
