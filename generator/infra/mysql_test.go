package infra

import (
	"testing"

	"github.com/rom8726/airo/config"

	"github.com/stretchr/testify/require"
)

func TestMySQLProcessor_AllMethods(t *testing.T) {
	p := NewMySQLProcessor()
	cfg := &config.ProjectConfig{ProjectName: "p", ModuleName: "m"}
	p.SetConfig(cfg)
	require.NotEmpty(t, p.Import())
	require.NotEmpty(t, p.Config())
	require.NotEmpty(t, p.ConfigField())
	require.NotEmpty(t, p.ConfigFieldName())
	require.NotEmpty(t, p.Constructor())
	require.NotEmpty(t, p.InitInAppConstructor())
	require.NotEmpty(t, p.StructField())
	require.NotEmpty(t, p.FillStructField())
	require.NotEmpty(t, p.Close())
	require.NotEmpty(t, p.DockerCompose())
	require.NotEmpty(t, p.ComposeEnv())
	require.NotEmpty(t, p.ConfigEnv())
	_ = p.MigrateFileData()
}
