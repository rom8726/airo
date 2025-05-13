package generator

import (
	"context"
	"fmt"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/generator/steps"
)

func GenerateProject(ctx context.Context, cfg *config.ProjectConfig) error {
	steps := []Step{
		steps.RootDirStep{},
		steps.GoModStep{},
		steps.SpecsStep{},
		steps.OGenStep{},
		steps.PkgStep{},
		steps.ConfigStep{},
		steps.RestAPIStep{},
		steps.AppStep{},
		steps.ServerCmdStep{},
		steps.MainGoStep{},
		steps.DockerfileStep{},
		steps.DevEnvStep{},
		steps.GolangCIStep{},
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
