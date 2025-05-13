package steps

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/rom8726/airo/config"
)

type GoModTidyStep struct{}

func (GoModTidyStep) Description() string {
	return "Run go mod tidy"
}

func (GoModTidyStep) Do(ctx context.Context, cfg *config.ProjectConfig) error {
	fmt.Println("")
	dir := projectDir(cfg)

	originalDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())

		return nil
	}
	defer func() {
		_ = os.Chdir(originalDir)
	}()

	if err := os.Chdir(dir); err != nil {
		fmt.Println(err.Error())

		return nil
	}

	cmd := exec.CommandContext(ctx, "go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
