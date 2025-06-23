package tui

import (
	"testing"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/generator/infra"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"
)

func newTestModel() *Model {
	reg := infra.NewRegistry()
	reg.RegisterDB("pg", "Postgres", nil, 1)
	cfg := &config.ProjectConfig{DB: "pg"}
	return InitialModel(cfg, reg)
}

func TestUpdate_ProjectNameToModuleName(t *testing.T) {
	m := newTestModel()
	m.input.SetValue("myproj")
	msg := tea.KeyMsg{Type: tea.KeyEnter}
	m2, _ := m.Update(msg)
	model := m2.(*Model)
	require.Equal(t, stepModuleName, model.step)
	require.Equal(t, "myproj", model.project)
}

func TestUpdate_ModuleNameToOpenAPIPath(t *testing.T) {
	m := newTestModel()
	m.step = stepModuleName
	m.input.SetValue("github.com/test/proj")
	msg := tea.KeyMsg{Type: tea.KeyEnter}
	m2, _ := m.Update(msg)
	model := m2.(*Model)
	require.Equal(t, stepOpenAPIPath, model.step)
	require.Equal(t, "github.com/test/proj", model.module)
}

func TestUpdate_Backspace_ReturnsToPrevStep(t *testing.T) {
	m := newTestModel()
	m.step = stepModuleName
	m.project = "myproj"
	m.input.SetValue("github.com/test/proj")
	msg := tea.KeyMsg{Type: tea.KeyBackspace}
	m2, _ := m.Update(msg)
	model := m2.(*Model)
	require.Equal(t, stepProjectName, model.step)
	require.Equal(t, "myproj", model.input.Value())
}

func TestUpdate_InvalidProjectName_ShowsError(t *testing.T) {
	m := newTestModel()
	m.input.SetValue("")
	msg := tea.KeyMsg{Type: tea.KeyEnter}
	m2, _ := m.Update(msg)
	model := m2.(*Model)
	require.Equal(t, stepProjectName, model.step)
	require.Empty(t, model.errMsg)
}

func TestUpdate_StepDone_FillsProjectConfig(t *testing.T) {
	m := newTestModel()
	m.step = stepDone
	m.project = "p"
	m.module = "m"
	m.openapiPath = "spec.yml"
	m.dbList.SetItems([]list.Item{dbItem{code: "pg", selected: true}})
	m.infraList.SetItems([]list.Item{infraItem{code: "redis", used: true}})
	m.testyList.SetItems([]list.Item{testyItem{selected: true}})
	msg := tea.KeyMsg{Type: tea.KeyEnter}
	cfg := &config.ProjectConfig{}
	m.projectConfig = cfg
	_, _ = m.Update(msg)
	require.Equal(t, "p", cfg.ProjectName)
	require.Equal(t, "m", cfg.ModuleName)
	require.Equal(t, "spec.yml", cfg.OpenAPIPath)
	require.Equal(t, "pg", cfg.DB)
	require.Equal(t, []string{"redis"}, cfg.UseInfra)
	require.True(t, cfg.UseTesty)
}
