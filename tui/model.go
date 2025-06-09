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
	stepTesty
	stepDone
)

const windowWidth = 80

type Model struct {
	step        step
	input       textinput.Model
	dbList      list.Model
	infraList   list.Model
	testyList   list.Model
	fileBrowser *FileBrowser
	confirmMsg  string
	errMsg      string
	errTS       time.Time

	project     string
	module      string
	infra       []infraItem
	openapiPath string

	projectConfig *config.ProjectConfig
}

func InitialModel(projectCfg *config.ProjectConfig, registry *infra.Registry) *Model {
	ti := textinput.New()
	ti.Placeholder = "Enter name"
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = windowWidth

	// Initialize file browser
	fb, err := NewFileBrowser("", windowWidth, 20)
	if err != nil {
		// If there's an error, we'll continue without the file browser
		// and fall back to text input for the file path
		fb = nil
	}

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
	dbList.Title = "Select a database ([space] to select)"
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
	infraList.Title = "Select infrastructure components ([space] to toggle)"
	infraList.SetShowStatusBar(false)
	infraList.SetFilteringEnabled(false)
	infraList.SetShowHelp(true)
	infraList.SetWidth(windowWidth)
	infraList.SetHeight(20)

	// ----- Testy -----
	testyItems := []list.Item{
		testyItem{
			title:    "Generate test files using Testy framework",
			selected: false,
		},
	}

	testyList := list.New(testyItems, list.NewDefaultDelegate(), 0, 0)
	testyList.Title = "Testing options ([space] to toggle)"
	testyList.SetShowStatusBar(false)
	testyList.SetFilteringEnabled(false)
	testyList.SetShowHelp(true)
	testyList.SetWidth(windowWidth)
	testyList.SetHeight(20)

	return &Model{
		projectConfig: projectCfg,
		step:          stepProjectName,
		input:         ti,
		dbList:        dbList,
		infraList:     infraList,
		testyList:     testyList,
		fileBrowser:   fb,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}
