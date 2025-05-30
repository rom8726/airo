package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/validate"
)

const errTimeout = 4 * time.Second

type clearErrMsg struct{}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.step {
	case stepProjectName, stepModuleName, stepOpenAPIPath:
		m.input, cmd = m.input.Update(msg)
	case stepDBChoice:
		m.dbList, cmd = m.dbList.Update(msg)
	case stepInfraChoice:
		m.infraList, cmd = m.infraList.Update(msg)
	case stepTesty:
		m.testyList, cmd = m.testyList.Update(msg)
	}

	switch msg := msg.(type) {
	case clearErrMsg:
		m.errMsg = ""

		return m, nil
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			*m.projectConfig = config.ProjectConfig{Aborted: true}

			return m, tea.Quit
		case tea.KeyEnter:
			switch m.step {
			case stepProjectName:
				m.project = m.input.Value()
				if m.project == "" {
					return m, nil
				}

				if err := validate.ValidateProjectName(m.project); err != nil {
					m.errMsg = err.Error()
					m.errTS = time.Now()

					return m, tea.Tick(errTimeout, func(time.Time) tea.Msg { return clearErrMsg{} })
				}

				m.input.SetValue("")
				m.input.Placeholder = "module name (e.g. github.com/user/myproject)"
				m.step = stepModuleName

			case stepModuleName:
				m.module = m.input.Value()
				if m.module == "" {
					return m, nil
				}

				if err := validate.ValidateModuleName(m.module); err != nil {
					m.errMsg = err.Error()
					m.errTS = time.Now()

					return m, tea.Tick(errTimeout, func(time.Time) tea.Msg { return clearErrMsg{} })
				}

				m.input.SetValue("")
				m.input.Placeholder = "Path to OpenAPI spec (e.g. example/server.yml)"
				m.step = stepOpenAPIPath

			case stepDBChoice:
				m.step = stepInfraChoice

			case stepInfraChoice:
				if getSelectedDB(m.dbList.Items()) == config.DBTypePostgres {
					m.step = stepTesty
				} else {
					m.step = stepDone
				}
			case stepTesty:
				m.step = stepDone

			case stepOpenAPIPath:
				m.openapiPath = m.input.Value()
				if m.openapiPath == "" {
					return m, nil
				}

				m.step = stepDBChoice

			case stepDone:
				*m.projectConfig = config.ProjectConfig{
					ProjectName: m.project,
					ModuleName:  m.module,
					OpenAPIPath: m.openapiPath,
					DB:          getSelectedDB(m.dbList.Items()),
					UseInfra:    getSelectedInfraCodes(m.infraList.Items()),
					UseTesty:    getSelectedTesty(m.testyList.Items()),
				}

				return m, tea.Quit
			}
		case tea.KeySpace:
			switch m.step {
			case stepDBChoice:
				i := m.dbList.Index()
				for idx, it := range m.dbList.Items() {
					if item, ok := it.(dbItem); ok {
						item.selected = idx == i
						m.dbList.SetItem(idx, item)
					}
				}
			case stepInfraChoice:
				i := m.infraList.Index()
				if item, ok := m.infraList.Items()[i].(infraItem); ok {
					item.used = !item.used
					m.infraList.SetItem(i, item)
				}
			case stepTesty:
				i := m.testyList.Index()
				if item, ok := m.testyList.Items()[i].(testyItem); ok {
					item.selected = !item.selected
					m.testyList.SetItem(i, item)
				}
			}
		}
	}

	if _, ok := msg.(tea.KeyMsg); ok && m.errMsg != "" {
		m.errMsg = ""
	}

	return m, cmd
}

func getSelectedDB(items []list.Item) string {
	for _, it := range items {
		if di, ok := it.(dbItem); ok && di.selected {
			return di.code
		}
	}

	return config.DBTypePostgres
}

func getSelectedInfraCodes(items []list.Item) []string {
	var result []string
	for _, it := range items {
		ii := it.(infraItem)
		if ii.used {
			result = append(result, ii.code)
		}
	}

	return result
}

func getSelectedTesty(items []list.Item) bool {
	for _, it := range items {
		if ti, ok := it.(testyItem); ok && ti.selected {
			return true
		}
	}

	return false
}
