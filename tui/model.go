package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/rom8726/airo/config"
)

type step int

const (
	stepProjectName step = iota
	stepModuleName
	stepInfraChoice
	stepOpenAPIPath
	stepDone
)

type infraItem struct {
	title string
	used  bool
}

func (i infraItem) Title() string {
	checked := "[ ]"
	if i.used {
		checked = "[x]"
	}
	return fmt.Sprintf("%s %s", checked, i.title)
}
func (i infraItem) Description() string { return "" }
func (i infraItem) FilterValue() string { return i.title }

type Model struct {
	step       step
	input      textinput.Model
	infraList  list.Model
	confirmMsg string

	project     string
	module      string
	infra       []infraItem
	openapiPath string

	projectConfig *config.ProjectConfig
}

func InitialModel(projectCfg *config.ProjectConfig) *Model {
	ti := textinput.New()
	ti.Placeholder = "project-name"
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 32

	items := []list.Item{
		infraItem{title: "Postgres"},
		infraItem{title: "Redis"},
	}

	infraList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	infraList.Title = "Choose infrastructure ([space] — make choose, [enter] — continue)"
	infraList.SetShowStatusBar(false)
	infraList.SetFilteringEnabled(false)
	infraList.SetShowHelp(true)
	infraList.SetWidth(80)
	infraList.SetHeight(12)

	return &Model{
		projectConfig: projectCfg,
		step:          stepProjectName,
		input:         ti,
		infraList:     infraList,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}
