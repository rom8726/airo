package generator

import (
	"context"
	"fmt"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/generator/infra"
	"github.com/rom8726/airo/generator/steps"
)

// StepProvider is a function that returns a Step
type StepProvider func(registry *infra.Registry) Step

// Generator handles the project generation process
type Generator struct {
	registry      *infra.Registry
	stepProviders []StepProvider
}

// New creates a new Generator with the given registry and default steps
func New(registry *infra.Registry) *Generator {
	return &Generator{
		registry:      registry,
		stepProviders: defaultStepProviders(),
	}
}

// WithSteps replaces the default steps with the provided steps
func (g *Generator) WithSteps(stepProviders []StepProvider) *Generator {
	g.stepProviders = stepProviders
	return g
}

// AddStep adds a step to the end of the steps list
func (g *Generator) AddStep(stepProvider StepProvider) *Generator {
	g.stepProviders = append(g.stepProviders, stepProvider)
	return g
}

// GenerateProject executes all steps to generate a project
func (g *Generator) GenerateProject(ctx context.Context, cfg *config.ProjectConfig) error {
	// Create steps from providers
	steps := make([]Step, 0, len(g.stepProviders))
	for _, provider := range g.stepProviders {
		steps = append(steps, provider(g.registry))
	}

	// Execute steps
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

// defaultStepProviders returns the default step providers
func defaultStepProviders() []StepProvider {
	return []StepProvider{
		func(r *infra.Registry) Step { return steps.RootDirStep{} },
		func(r *infra.Registry) Step { return steps.GoModStep{} },
		func(r *infra.Registry) Step { return steps.SpecsStep{} },
		func(r *infra.Registry) Step { return steps.OGenStep{} },
		func(r *infra.Registry) Step { return steps.PkgStep{} },
		func(r *infra.Registry) Step { return steps.NewConfigStep(r) },
		func(r *infra.Registry) Step { return steps.RestAPIStep{} },
		func(r *infra.Registry) Step { return steps.NewAppStep(r) },
		func(r *infra.Registry) Step { return steps.NewServerCmdStep(r) },
		func(r *infra.Registry) Step { return steps.MainGoStep{} },
		func(r *infra.Registry) Step { return steps.NewMigrateStep(r) },
		func(r *infra.Registry) Step { return steps.DockerfileStep{} },
		func(r *infra.Registry) Step { return steps.NewDevEnvStep(r) },
		func(r *infra.Registry) Step { return steps.GolangCIStep{} },
		func(r *infra.Registry) Step { return steps.RealtimeStep{} },
		func(r *infra.Registry) Step { return steps.TestyStep{} },
		func(r *infra.Registry) Step { return steps.GoModTidyStep{} },
	}
}
