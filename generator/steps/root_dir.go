package steps

import (
	"context"
	"os"

	"github.com/rom8726/airo/config"
)

type RootDirStep struct{}

func (RootDirStep) Description() string {
	return "Create project directory"
}

func (RootDirStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
	root := projectDir(cfg)
	if err := os.MkdirAll(root, 0755); err != nil {
		return err
	}

	return nil
}
