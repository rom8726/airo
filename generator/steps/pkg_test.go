package steps

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/rom8726/airo/config"
)

func TestPkgStep_Do_Success(t *testing.T) {
	dir := t.TempDir()
	cfg := &config.ProjectConfig{ProjectName: dir, DB: config.DBTypePostgres}
	step := PkgStep{}
	err := step.Do(context.Background(), cfg)
	require.NoError(t, err)
	path := filepath.Join(dir, "pkg", "httpserver", "httpserver.go")
	_, err = os.Stat(path)
	require.NoError(t, err)
}

func TestPkgStep_Do_MkdirError(t *testing.T) {
	cfg := &config.ProjectConfig{ProjectName: "/dev/null/forbidden", DB: config.DBTypePostgres}
	step := PkgStep{}
	err := step.Do(context.Background(), cfg)
	require.Error(t, err)
}
