package steps

import (
	"path/filepath"

	"github.com/rom8726/airo/config"
)

func projectDir(cfg *config.ProjectConfig) string {
	return cfg.ProjectName
}

func openapiDir(cfg *config.ProjectConfig) string {
	return filepath.Join(projectDir(cfg), "internal", "generated", "server")
}

func configDir(cfg *config.ProjectConfig) string {
	return filepath.Join(projectDir(cfg), "internal", "config")
}
