package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/rom8726/airo/config"
)

type step int

const (
	stepProjectName step = iota
	stepModuleName
	stepOpenAPIPath
	stepDBChoice
	stepInfraChoice
	stepDone
)

const (
	postgresName = "PostgreSQL"
	mysqlName    = "MySQL"
	redisName    = "Redis"
	kafkaName    = "Kafka"
)

type Model struct {
	step       step
	input      textinput.Model
	dbList     list.Model
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

	// ----- DB -----
	dbItems := []list.Item{
		dbItem{title: postgresName, selected: true},
		dbItem{title: mysqlName},
	}
	dbList := list.New(dbItems, list.NewDefaultDelegate(), 0, 0)
	dbList.Title = "Choose database ([space] — select, [enter] — continue)"
	dbList.SetShowStatusBar(false)
	dbList.SetFilteringEnabled(false)
	dbList.SetShowHelp(true)
	dbList.SetWidth(80)
	dbList.SetHeight(12)

	// ----- Infra -----
	items := []list.Item{
		infraItem{title: redisName},
		infraItem{title: kafkaName},
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
		dbList:        dbList,
		infraList:     infraList,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}
