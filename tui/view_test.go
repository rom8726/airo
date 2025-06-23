package tui

import (
	"testing"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/generator/infra"

	"github.com/charmbracelet/bubbles/list"
	"github.com/stretchr/testify/require"
)

func newViewTestModel() *Model {
	reg := infra.NewRegistry()
	reg.RegisterDB("pg", "Postgres", nil, 1)
	reg.RegisterInfra("redis", "Redis", nil, 1)
	cfg := &config.ProjectConfig{DB: "pg"}
	return InitialModel(cfg, reg)
}

func TestView_Steps(t *testing.T) {
	m := newViewTestModel()
	for s := stepProjectName; s <= stepDone; s++ {
		m.step = s
		out := m.View()
		require.NotEmpty(t, out)
	}
}

func TestView_ErrorMessage(t *testing.T) {
	m := newViewTestModel()
	m.step = stepModuleName
	m.errMsg = "some error"
	out := m.View()
	require.Contains(t, out, "some error")
}

func TestView_Done_LongOpenAPIPathAndInfra(t *testing.T) {
	t.SkipNow()

	m := newViewTestModel()
	m.step = stepDone
	m.project = "proj"
	m.module = "mod"
	m.openapiPath = "/very/long/path/to/openapi/specification/file/that/is/way/too/long/openapi.yml"
	infraCodes := []string{"redis", "kafka", "nats", "rabbitmq", "elasticsearch", "memcache", "etcd", "aerospike", "mongo"}
	items := make([]list.Item, len(infraCodes))
	for i, code := range infraCodes {
		items[i] = infraItem{code: code, title: code, used: true}
	}
	m.infraList.SetItems(items)
	out := m.View()
	require.Contains(t, out, "...openapi.yml")
	require.Contains(t, out, "Infra:    ")
}

func TestView_Done_AllFieldsFilled(t *testing.T) {
	m := newViewTestModel()
	m.step = stepDone
	m.project = "proj"
	m.module = "mod"
	m.openapiPath = "api.yml"
	m.dbList.SetItems([]list.Item{dbItem{code: "pg", selected: true}})
	m.infraList.SetItems([]list.Item{infraItem{code: "redis", used: true}})
	m.testyList.SetItems([]list.Item{testyItem{selected: true}})
	out := m.View()
	require.Contains(t, out, "Project:  proj")
	require.Contains(t, out, "Module:   mod")
	require.Contains(t, out, "OpenAPI:  api.yml")
	require.Contains(t, out, "Database: pg")
	require.Contains(t, out, "Infra:    redis")
	require.Contains(t, out, "Testy:    Yes")
}

func TestView_Done_NoInfraNoTesty(t *testing.T) {
	m := newViewTestModel()
	m.step = stepDone
	m.project = "proj"
	m.module = "mod"
	m.openapiPath = "api.yml"
	m.dbList.SetItems([]list.Item{dbItem{code: "pg", selected: true}})
	m.infraList.SetItems([]list.Item{})
	m.testyList.SetItems([]list.Item{testyItem{selected: false}})
	out := m.View()
	require.Contains(t, out, "Infra:    None")
	require.Contains(t, out, "Testy:    No")
}
