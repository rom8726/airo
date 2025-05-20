package steps

import (
	"context"
	"fmt"
	"os"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/generator/infra"
)

type MigrateStep struct {
	registry *infra.Registry
}

func NewMigrateStep(registry *infra.Registry) *MigrateStep {
	return &MigrateStep{
		registry: registry,
	}
}

func (s *MigrateStep) Description() string {
	return "Add migrations functionality"
}

func (s *MigrateStep) Do(_ context.Context, cfg *config.ProjectConfig) error {
	dir := migrationsDir(cfg)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}

	db := s.registry.GetDB(cfg.DB)
	migrateFileData := db.Processor.MigrateFileData()
	if len(migrateFileData) == 0 {
		return nil
	}

	migrateFilePath := migrateGoPath(cfg)
	err := os.WriteFile(migrateFilePath, migrateFileData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write migrate.go: %w", err)
	}

	return nil
}
