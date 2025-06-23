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

func TestDevEnvStep_Do_Success(t *testing.T) {
	reg := infra.NewRegistry()
	reg.RegisterDB(config.DBTypePostgres, "Postgres", &mockProcessor{}, 1)
	dir := t.TempDir()
	cfg := &config.ProjectConfig{ProjectName: dir, ModuleName: "mod", DB: config.DBTypePostgres}
	step := NewDevEnvStep(reg)
	err := step.Do(context.Background(), cfg)
	require.NoError(t, err)
	makefilePath := filepath.Join(dir, "Makefile")
	_, err = os.Stat(makefilePath)
	require.NoError(t, err)
	composePath := filepath.Join(dir, "dev", "docker-compose.yml")
	_, err = os.Stat(composePath)
	require.NoError(t, err)
}

func TestDevEnvStep_Do_MkdirError(t *testing.T) {
	reg := infra.NewRegistry()
	cfg := &config.ProjectConfig{ProjectName: "/dev/null/forbidden", ModuleName: "mod", DB: config.DBTypePostgres}
	step := NewDevEnvStep(reg)
	err := step.Do(context.Background(), cfg)
	require.Error(t, err)
}
