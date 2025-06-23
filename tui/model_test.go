package tui

import (
	"testing"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/generator/infra"

	"github.com/stretchr/testify/require"
)

func TestInitialModel_Basic(t *testing.T) {
	reg := infra.NewRegistry()
	reg.RegisterDB("pg", "Postgres", nil, 1)
	reg.RegisterInfra("redis", "Redis", nil, 1)
	cfg := &config.ProjectConfig{
		ProjectName: "testproj",
		ModuleName:  "github.com/test/proj",
		DB:          "pg",
		UseInfra:    []string{"redis"},
	}
	model := InitialModel(cfg, reg)
	require.NotNil(t, model)
	require.Equal(t, stepProjectName, model.step)
	require.Equal(t, cfg, model.projectConfig)
	require.NotNil(t, model.input)
	require.NotNil(t, model.dbList)
	require.NotNil(t, model.infraList)
	require.NotNil(t, model.testyList)
	// fileBrowser может быть nil, если getwd не сработал, но обычно не nil
}

func TestInitialModel_DBList(t *testing.T) {
	reg := infra.NewRegistry()
	reg.RegisterDB("pg", "Postgres", nil, 1)
	reg.RegisterDB("my", "MySQL", nil, 2)
	cfg := &config.ProjectConfig{DB: "pg"}
	model := InitialModel(cfg, reg)
	dbItems := model.dbList.Items()
	require.Len(t, dbItems, 2)
	found := false
	for _, it := range dbItems {
		if di, ok := it.(dbItem); ok && di.code == "pg" {
			found = true
		}
	}
	require.True(t, found)
}

func TestInitialModel_InfraList(t *testing.T) {
	reg := infra.NewRegistry()
	reg.RegisterInfra("redis", "Redis", nil, 1)
	reg.RegisterInfra("kafka", "Kafka", nil, 2)
	cfg := &config.ProjectConfig{}
	model := InitialModel(cfg, reg)
	infraItems := model.infraList.Items()
	require.Len(t, infraItems, 2)
}

func TestModel_Init(t *testing.T) {
	reg := infra.NewRegistry()
	reg.RegisterDB("pg", "Postgres", nil, 1)
	cfg := &config.ProjectConfig{DB: "pg"}
	model := InitialModel(cfg, reg)
	cmd := model.Init()
	require.Nil(t, cmd)
}
