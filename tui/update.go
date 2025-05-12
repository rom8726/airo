package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/rom8726/airo/config"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.step {
	case stepProjectName, stepModuleName, stepOpenAPIPath:
		m.input, cmd = m.input.Update(msg)
	case stepDBChoice:
		m.dbList, cmd = m.dbList.Update(msg)
	case stepInfraChoice:
		m.infraList, cmd = m.infraList.Update(msg)
	}

	switch msg := msg.(type) {
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

				m.input.SetValue("")
				m.input.Placeholder = "module name (e.g. github.com/user/project)"
				m.step = stepModuleName

			case stepModuleName:
				m.module = m.input.Value()
				if m.module == "" {
					return m, nil
				}

				m.input.SetValue("")
				m.input.Placeholder = "Path to OpenAPI spec"
				m.step = stepOpenAPIPath

			case stepDBChoice:
				m.step = stepInfraChoice

			case stepInfraChoice:
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
					DB:          dbType(getSelectedDB(m.dbList.Items())),
					UseInfra:    getSelectedInfraCodes(m.infraList.Items()),
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
			}
		}
	}

	return m, cmd
}

func getSelectedDB(items []list.Item) string {
	for _, it := range items {
		if di, ok := it.(dbItem); ok && di.selected {
			return di.title
		}
	}

	return postgresName
}

func dbType(selectedDB string) config.DBType {
	switch selectedDB {
	case postgresName:
		return config.DBTypePostgres
	case mysqlName:
		return config.DBTypeMySQL
	default:
		return config.DBTypeUnknown
	}
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

func contains(list []string, elem string) bool {
	for _, a := range list {
		if a == elem {
			return true
		}
	}

	return false
}
