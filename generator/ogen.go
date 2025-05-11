package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/rom8726/airo/config"
)

func runOGen(cfg *config.ProjectConfig) error {
	currDir := os.Getenv("PWD")

	outputDir := filepath.Join(cfg.ProjectName, "internal", "generated", "server")
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}

	specPath := cfg.OpenAPIPath
	targetSpecPath := filepath.Join("/workspace", specPath)
	targetOutputDir := filepath.Join("/workspace", outputDir)

	cmd := exec.Command(
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

	fmt.Println(cmd.String())

	return cmd.Run()
}
