package steps

import (
	"context"
	"fmt"
	"os"

	"github.com/rom8726/airo/config"
)

type SpecsStep struct{}

func (SpecsStep) Description() string {
	return "Copy OpenAPI specs"
}

func (SpecsStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
	dir := specsDir(cfg)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}

	src := cfg.OpenAPIPath
	target := serverSpecPath(cfg)

	content, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("failed to read source file: %w", err)
	}

	if err := os.WriteFile(target, content, 0644); err != nil {
		return fmt.Errorf("failed to write target file: %w", err)
	}

	return nil
}
