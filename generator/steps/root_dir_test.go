package steps

import (
	"context"
	"os"
	"testing"

	"github.com/rom8726/airo/config"

	"github.com/stretchr/testify/require"
)

func TestRootDirStep_Do_Success(t *testing.T) {
	dir := t.TempDir()
	cfg := &config.ProjectConfig{ProjectName: dir}
	step := RootDirStep{}
	err := step.Do(context.Background(), cfg)
	require.NoError(t, err)
	_, err = os.Stat(dir)
	require.NoError(t, err)
}

func TestRootDirStep_Do_MkdirError(t *testing.T) {
	cfg := &config.ProjectConfig{ProjectName: "/dev/null/forbidden"}
	step := RootDirStep{}
	err := step.Do(context.Background(), cfg)
	require.Error(t, err)
}
