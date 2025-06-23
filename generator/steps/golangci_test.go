package steps

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/rom8726/airo/config"

	"github.com/stretchr/testify/require"
)

func TestGolangCIStep_Do_Success(t *testing.T) {
	dir := t.TempDir()
	cfg := &config.ProjectConfig{ProjectName: dir}
	step := GolangCIStep{}
	err := step.Do(context.Background(), cfg)
	require.NoError(t, err)
	path := filepath.Join(dir, ".golangci.yml")
	_, err = os.Stat(path)
	require.NoError(t, err)
}

func TestGolangCIStep_Do_MkdirError(t *testing.T) {
	cfg := &config.ProjectConfig{ProjectName: "/dev/null/forbidden"}
	step := GolangCIStep{}
	err := step.Do(context.Background(), cfg)
	require.Error(t, err)
}
