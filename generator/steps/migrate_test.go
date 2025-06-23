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

func TestMigrateStep_Do_Success(t *testing.T) {
	reg := infra.NewRegistry()
	reg.RegisterDB(config.DBTypePostgres, "Postgres", &mockProcessorWithMigrateFile{data: []byte("test")}, 1)
	dir := t.TempDir()
	// Ensure a cmd / server directory exists
	os.MkdirAll(filepath.Join(dir, "cmd", "server"), 0755)
	cfg := &config.ProjectConfig{ProjectName: dir, DB: config.DBTypePostgres}
	step := NewMigrateStep(reg)
	err := step.Do(context.Background(), cfg)
	require.NoError(t, err)
	path := filepath.Join(dir, "cmd", "server", "migrate.go")
	_, err = os.Stat(path)
	require.NoError(t, err)
}

type mockProcessorWithMigrateFile struct{ data []byte }

func (m *mockProcessorWithMigrateFile) SetConfig(cfg *config.ProjectConfig) {}
func (m *mockProcessorWithMigrateFile) Import() string                      { return "" }
func (m *mockProcessorWithMigrateFile) Config() string                      { return "" }
func (m *mockProcessorWithMigrateFile) ConfigField() string                 { return "" }
func (m *mockProcessorWithMigrateFile) ConfigFieldName() string             { return "" }
func (m *mockProcessorWithMigrateFile) Constructor() string                 { return "" }
func (m *mockProcessorWithMigrateFile) InitInAppConstructor() string        { return "" }
func (m *mockProcessorWithMigrateFile) StructField() string                 { return "" }
func (m *mockProcessorWithMigrateFile) FillStructField() string             { return "" }
func (m *mockProcessorWithMigrateFile) Close() string                       { return "" }
func (m *mockProcessorWithMigrateFile) DockerCompose() string               { return "" }
func (m *mockProcessorWithMigrateFile) ComposeEnv() string                  { return "" }
func (m *mockProcessorWithMigrateFile) ConfigEnv() string                   { return "" }
func (m *mockProcessorWithMigrateFile) MigrateFileData() []byte             { return m.data }

func TestMigrateStep_Do_MkdirError(t *testing.T) {
	reg := infra.NewRegistry()
	reg.RegisterDB(config.DBTypePostgres, "Postgres", &mockProcessorWithMigrateFile{data: []byte("test")}, 1)
	cfg := &config.ProjectConfig{ProjectName: "/dev/null/forbidden", DB: config.DBTypePostgres}
	step := NewMigrateStep(reg)
	err := step.Do(context.Background(), cfg)
	require.Error(t, err)
}
