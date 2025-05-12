package steps

import (
	"path/filepath"

	"github.com/rom8726/airo/config"
)

func projectDir(cfg *config.ProjectConfig) string {
	return cfg.ProjectName
}

func internalDir(cfg *config.ProjectConfig) string {
	return filepath.Join(projectDir(cfg), "internal")
}

func openapiDir(cfg *config.ProjectConfig) string {
	return filepath.Join(projectDir(cfg), "internal", "generated", "server")
}

func configDir(cfg *config.ProjectConfig) string {
	return filepath.Join(projectDir(cfg), "internal", "config")
}

func serverCmdDir(cfg *config.ProjectConfig) string {
	return filepath.Join(projectDir(cfg), "cmd", "server")
}

func restAPIDir(cfg *config.ProjectConfig) string {
	return filepath.Join(projectDir(cfg), "internal", "api", "rest")
}

func pkgDir(cfg *config.ProjectConfig) string {
	return filepath.Join(projectDir(cfg), "pkg")
}

func pkgHttpServerDir(cfg *config.ProjectConfig) string {
	return filepath.Join(pkgDir(cfg), "httpserver")
}

func specsDir(cfg *config.ProjectConfig) string {
	return filepath.Join(projectDir(cfg), "specs")
}

func serverSpecPath(cfg *config.ProjectConfig) string {
	return filepath.Join(specsDir(cfg), "server.yml")
}
