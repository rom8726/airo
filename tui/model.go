package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/generator/infra"
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

const windowWidth = 80

type Model struct {
	step       step
	input      textinput.Model
	dbList     list.Model
	infraList  list.Model
	confirmMsg string
	errMsg     string
	errTS      time.Time

	project     string
	module      string
	infra       []infraItem
	openapiPath string

	projectConfig *config.ProjectConfig
}

func InitialModel(projectCfg *config.ProjectConfig, registry *infra.Registry) *Model {
	ti := textinput.New()
	ti.Placeholder = "project-name"
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = windowWidth

	// ----- DB -----
	dbInfos := registry.ListDBs()
	dbItems := make([]list.Item, 0, len(dbInfos))
	for _, elem := range dbInfos {
		dbItems = append(dbItems, dbItem{
			title:    elem.Title,
			code:     elem.Code,
			selected: elem.Code == config.DBTypePostgres,
		})
	}

	dbList := list.New(dbItems, list.NewDefaultDelegate(), 0, 0)
	dbList.Title = "Choose database ([space] — select, [enter] — continue)"
	dbList.SetShowStatusBar(false)
	dbList.SetFilteringEnabled(false)
	dbList.SetShowHelp(true)
	dbList.SetWidth(windowWidth)
	dbList.SetHeight(20)

	// ----- Infra -----
	infraInfos := registry.ListInfras()
	items := make([]list.Item, 0, len(infraInfos))
	for _, elem := range infraInfos {
		items = append(items, infraItem{
			title: elem.Title,
			code:  elem.Code,
		})
	}

	infraList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	infraList.Title = "Choose infrastructure ([space] — make choose, [enter] — continue)"
	infraList.SetShowStatusBar(false)
	infraList.SetFilteringEnabled(false)
	infraList.SetShowHelp(true)
	infraList.SetWidth(windowWidth)
	infraList.SetHeight(20)

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
