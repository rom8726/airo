package steps

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/rom8726/airo/config"

	"github.com/stretchr/testify/require"
)

func TestRestAPIStep_Do_Success(t *testing.T) {
	dir := t.TempDir()
	cfg := &config.ProjectConfig{ProjectName: dir, ModuleName: "mod"}
	step := RestAPIStep{}

	// Мокаем шаблон
	oldTmpl := tmplApiGo
	tmplApiGo = "package test\n// {{.Module}}"
	defer func() { tmplApiGo = oldTmpl }()

	err := step.Do(context.Background(), cfg)
	require.NoError(t, err)
	path := filepath.Join(dir, "internal", "api", "rest", "api.go")
	_, err = os.Stat(path)
	require.NoError(t, err)
}

func TestRestAPIStep_Do_MkdirError(t *testing.T) {
	cfg := &config.ProjectConfig{ProjectName: "/dev/null/forbidden", ModuleName: "mod"}
	step := RestAPIStep{}

	// Мокаем шаблон
	oldTmpl := tmplApiGo
	tmplApiGo = "package test\n// {{.Module}}"
	defer func() { tmplApiGo = oldTmpl }()

	err := step.Do(context.Background(), cfg)
	require.Error(t, err)
}
