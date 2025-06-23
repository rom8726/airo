package steps

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/rom8726/airo/config"

	"github.com/stretchr/testify/require"
)

func TestTestyStep_Do_Success(t *testing.T) {
	dir := t.TempDir()
	cfg := &config.ProjectConfig{ProjectName: dir, ModuleName: "mod", DB: config.DBTypePostgres, UseTesty: true}
	step := TestyStep{}
	err := step.Do(context.Background(), cfg)
	require.NoError(t, err)
	// Проверяем, что создался env.go
	path := filepath.Join(dir, "tests", "runner", "env.go")
	_, err = os.Stat(path)
	require.NoError(t, err)
}

func TestTestyStep_Do_SkipIfNotEnabled(t *testing.T) {
	dir := t.TempDir()
	cfg := &config.ProjectConfig{ProjectName: dir, UseTesty: false}
	step := TestyStep{}
	err := step.Do(context.Background(), cfg)
	require.NoError(t, err)
}

func TestTestyStep_Do_MkdirError(t *testing.T) {
	cfg := &config.ProjectConfig{ProjectName: "/dev/null/forbidden", UseTesty: true}
	step := TestyStep{}
	err := step.Do(context.Background(), cfg)
	require.Error(t, err)
}
