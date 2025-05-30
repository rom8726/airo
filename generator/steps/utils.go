package steps

import (
	"os"
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
	return filepath.Join(projectDir(cfg), openapiRelDir())
}

func openapiRelDir() string {
	return filepath.Join("internal", "generated", "server")
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

func pkgDBDir(cfg *config.ProjectConfig) string {
	return filepath.Join(pkgDir(cfg), "db")
}

func pkgKafkaDir(cfg *config.ProjectConfig) string {
	return filepath.Join(pkgDir(cfg), "kafka")
}

func specsDir(cfg *config.ProjectConfig) string {
	return filepath.Join(projectDir(cfg), "specs")
}

func serverSpecPath(cfg *config.ProjectConfig) string {
	return filepath.Join(specsDir(cfg), "server.yml")
}

func securityHandlerPath(cfg *config.ProjectConfig) string {
	return filepath.Join(openapiDir(cfg), securityHandlerFileName)
}

func devDir(cfg *config.ProjectConfig) string {
	return filepath.Join(projectDir(cfg), "dev")
}

func migrationsDir(cfg *config.ProjectConfig) string {
	return filepath.Join(projectDir(cfg), "migrations")
}

func migrateGoPath(cfg *config.ProjectConfig) string {
	return filepath.Join(serverCmdDir(cfg), "migrate.go")
}

func hasSecurityHandler(cfg *config.ProjectConfig) bool {
	_, err := os.Stat(securityHandlerPath(cfg))

	return !os.IsNotExist(err)
}
