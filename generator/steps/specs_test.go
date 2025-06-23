package steps

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/rom8726/airo/config"

	"github.com/stretchr/testify/require"
)

func TestSpecsStep_Do_Success(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "openapi.yml")
	os.WriteFile(src, []byte("openapi: 3.0.0\ninfo:\n  title: Test\n"), 0644)
	cfg := &config.ProjectConfig{ProjectName: dir, OpenAPIPath: src}
	step := SpecsStep{}
	err := step.Do(context.Background(), cfg)
	require.NoError(t, err)
	path := filepath.Join(dir, "specs", "server.yml")
	_, err = os.Stat(path)
	require.NoError(t, err)
}

func TestSpecsStep_Do_MkdirError(t *testing.T) {
	cfg := &config.ProjectConfig{ProjectName: "/dev/null/forbidden", OpenAPIPath: "/dev/null/forbidden/openapi.yml"}
	step := SpecsStep{}
	err := step.Do(context.Background(), cfg)
	require.Error(t, err)
}
