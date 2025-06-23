package steps

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/rom8726/airo/config"

	"github.com/stretchr/testify/require"
)

func TestUtils_Paths(t *testing.T) {
	cfg := &config.ProjectConfig{ProjectName: "proj", ModuleName: "mod"}
	require.Equal(t, "proj", projectDir(cfg))
	require.Contains(t, internalDir(cfg), filepath.Join("proj", "internal"))
	require.Contains(t, openapiDir(cfg), filepath.Join("proj", "internal", "generated", "server"))
	require.Contains(t, configDir(cfg), filepath.Join("proj", "internal", "config"))
	require.Contains(t, serverCmdDir(cfg), filepath.Join("proj", "cmd", "server"))
	require.Contains(t, restAPIDir(cfg), filepath.Join("proj", "internal", "api", "rest"))
	require.Contains(t, pkgDir(cfg), filepath.Join("proj", "pkg"))
	require.Contains(t, pkgHttpServerDir(cfg), filepath.Join("proj", "pkg", "httpserver"))
	require.Contains(t, pkgDBDir(cfg), filepath.Join("proj", "pkg", "db"))
	require.Contains(t, pkgKafkaDir(cfg), filepath.Join("proj", "pkg", "kafka"))
	require.Contains(t, specsDir(cfg), filepath.Join("proj", "specs"))
	require.Contains(t, serverSpecPath(cfg), filepath.Join("proj", "specs", "server.yml"))
	require.Contains(t, securityHandlerPath(cfg), filepath.Join("proj", "internal", "generated", "server", "oas_security_gen.go"))
	require.Contains(t, devDir(cfg), filepath.Join("proj", "dev"))
	require.Contains(t, migrationsDir(cfg), filepath.Join("proj", "migrations"))
	require.Contains(t, migrateGoPath(cfg), filepath.Join("proj", "cmd", "server", "migrate.go"))
	require.Contains(t, testsRunnerDir(cfg), filepath.Join("proj", "tests", "runner"))
}

func TestUtils_HasSecurityHandler(t *testing.T) {
	dir := t.TempDir()
	cfg := &config.ProjectConfig{ProjectName: dir, ModuleName: "mod"}
	path := securityHandlerPath(cfg)
	os.MkdirAll(filepath.Dir(path), 0755)
	_, err := os.Stat(path)
	require.True(t, os.IsNotExist(err))
	require.False(t, hasSecurityHandler(cfg))
	os.WriteFile(path, []byte("test"), 0644)
	require.True(t, hasSecurityHandler(cfg))
}
