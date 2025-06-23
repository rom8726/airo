package steps

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/generator/infra"

	"github.com/stretchr/testify/require"
)

func TestAppStep_Do_Success(t *testing.T) {
	reg := infra.NewRegistry()
	// Register a mock DB processor
	reg.RegisterDB(config.DBTypePostgres, "Postgres", &mockProcessor{}, 1)
	dir := t.TempDir()
	cfg := &config.ProjectConfig{ProjectName: dir, ModuleName: "mod", DB: config.DBTypePostgres}
	step := NewAppStep(reg)
	err := step.Do(context.Background(), cfg)
	require.NoError(t, err)
	path := filepath.Join(dir, "internal", "app.go")
	_, err = os.Stat(path)
	require.NoError(t, err)
}

func TestAppStep_Do_MkdirError(t *testing.T) {
	reg := infra.NewRegistry()
	cfg := &config.ProjectConfig{ProjectName: "/dev/null/forbidden", ModuleName: "mod", DB: config.DBTypePostgres}
	step := NewAppStep(reg)
	err := step.Do(context.Background(), cfg)
	require.Error(t, err)
}
