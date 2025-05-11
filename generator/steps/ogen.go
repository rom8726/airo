package steps

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/rom8726/airo/config"
)

type OGenStep struct{}

func (OGenStep) Description() string {
	return "Generate OpenAPI server code"
}

func (OGenStep) Do(ctx context.Context, cfg *config.ProjectConfig) error {
	currDir := os.Getenv("PWD")

	outputDir := openapiDir(cfg)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}

	specPath := cfg.OpenAPIPath
	targetSpecPath := filepath.Join("/workspace", specPath)
	targetOutputDir := filepath.Join("/workspace", outputDir)

	cmd := exec.CommandContext(
		ctx,
		"docker",
		"run",
		"--rm",
		"-v", currDir+":/workspace",
		ogenDockerImage,
		"--target", targetOutputDir,
		"--clean",
		targetSpecPath,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("\n%s\n", cmd.String())

	return cmd.Run()
}
