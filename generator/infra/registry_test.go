package infra

import (
	"testing"

	"github.com/rom8726/airo/config"

	"github.com/stretchr/testify/require"
)

type mockProcessor struct {
	setConfigCalled bool
	cfg             *config.ProjectConfig
}

func (m *mockProcessor) SetConfig(cfg *config.ProjectConfig) { m.setConfigCalled = true; m.cfg = cfg }
func (m *mockProcessor) Import() string                      { return "import" }
func (m *mockProcessor) Config() string                      { return "config" }
func (m *mockProcessor) ConfigField() string                 { return "field" }
func (m *mockProcessor) ConfigFieldName() string             { return "fieldName" }
func (m *mockProcessor) Constructor() string                 { return "ctor" }
func (m *mockProcessor) InitInAppConstructor() string        { return "init" }
func (m *mockProcessor) StructField() string                 { return "structField" }
func (m *mockProcessor) FillStructField() string             { return "fill" }
func (m *mockProcessor) Close() string                       { return "close" }
func (m *mockProcessor) DockerCompose() string               { return "docker" }
func (m *mockProcessor) ComposeEnv() string                  { return "composeEnv" }
func (m *mockProcessor) ConfigEnv() string                   { return "configEnv" }
func (m *mockProcessor) MigrateFileData() []byte             { return nil }

func TestRegistry_AddAndGetDB(t *testing.T) {
	reg := NewRegistry()
	proc := &mockProcessor{}
	reg.RegisterDB("pg", "Postgres", proc, 2)
	reg.RegisterDB("my", "MySQL", proc, 1)

	dbs := reg.ListDBs()
	require.Len(t, dbs, 2)
	require.Equal(t, "MySQL", dbs[0].Title)
	db := reg.GetDB("pg")
	require.Equal(t, "Postgres", db.Title)
}

func TestRegistry_AddAndGetInfra(t *testing.T) {
	reg := NewRegistry()
	proc := &mockProcessor{}
	reg.RegisterInfra("redis", "Redis", proc, 1)
	reg.RegisterInfra("kafka", "Kafka", proc, 2)

	infras := reg.ListInfras()
	require.Len(t, infras, 2)
	require.Equal(t, "Redis", infras[0].Title)
	infra := reg.GetInfra("kafka")
	require.Equal(t, "Kafka", infra.Title)
}

func TestRegistry_UpdateConfig(t *testing.T) {
	reg := NewRegistry()
	proc1 := &mockProcessor{}
	proc2 := &mockProcessor{}
	reg.RegisterDB("pg", "Postgres", proc1, 1)
	reg.RegisterInfra("redis", "Redis", proc2, 1)
	cfg := &config.ProjectConfig{ProjectName: "test"}

	reg.UpdateConfig(cfg)
	require.True(t, proc1.setConfigCalled)
	require.True(t, proc2.setConfigCalled)
	require.Equal(t, cfg, proc1.cfg)
	require.Equal(t, cfg, proc2.cfg)
}
