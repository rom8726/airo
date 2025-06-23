package steps

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/rom8726/airo/config"

	"github.com/stretchr/testify/require"
)

func TestGoModStep_Do_Success(t *testing.T) {
	dir := t.TempDir()
	cfg := &config.ProjectConfig{ProjectName: dir, ModuleName: "mod"}
	step := GoModStep{}
	err := step.Do(context.Background(), cfg)
	require.NoError(t, err)
	path := filepath.Join(dir, "go.mod")
	_, err = os.Stat(path)
	require.NoError(t, err)
}

func TestGoModStep_Do_MkdirError(t *testing.T) {
	cfg := &config.ProjectConfig{ProjectName: "/dev/null/forbidden", ModuleName: "mod"}
	step := GoModStep{}
	err := step.Do(context.Background(), cfg)
	require.Error(t, err)
}
