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

func TestServerCmdStep_Do_Success(t *testing.T) {
	reg := infra.NewRegistry()
	reg.RegisterDB(config.DBTypePostgres, "Postgres", &mockProcessor{}, 1)
	dir := t.TempDir()
	os.MkdirAll(filepath.Join(dir, "cmd", "server"), 0755)
	cfg := &config.ProjectConfig{ProjectName: dir, ModuleName: "mod", DB: config.DBTypePostgres}
	step := NewServerCmdStep(reg)
	err := step.Do(context.Background(), cfg)
	if err != nil {
		t.Logf("step.Do error: %v", err)
	}
	require.NoError(t, err)
	path := filepath.Join(dir, "cmd", "server", "server.go")
	_, err = os.Stat(path)
	if err != nil {
		t.Logf("os.Stat error: %v", err)
	}
	require.NoError(t, err)
}

func TestServerCmdStep_Do_MkdirError(t *testing.T) {
	reg := infra.NewRegistry()
	cfg := &config.ProjectConfig{ProjectName: "/dev/null/forbidden", ModuleName: "mod", DB: config.DBTypePostgres}
	step := NewServerCmdStep(reg)
	err := step.Do(context.Background(), cfg)
	require.Error(t, err)
}
