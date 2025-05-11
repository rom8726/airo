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
	case stepInfraChoice:
		m.infraList, cmd = m.infraList.Update(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			switch m.step {
			case stepProjectName:
				m.project = m.input.Value()
				m.input.SetValue("")
				m.input.Placeholder = "module name (e.g. github.com/user/project)"
				m.step = stepModuleName

			case stepModuleName:
				m.module = m.input.Value()
				m.step = stepInfraChoice

			case stepInfraChoice:
				var infras []string
				for _, it := range m.infraList.Items() {
					ii := it.(infraItem)
					if ii.used {
						infras = append(infras, ii.title)
					}
				}

				m.input.SetValue("")
				m.input.Placeholder = "Path to OpenAPI spec"
				m.step = stepOpenAPIPath

			case stepOpenAPIPath:
				m.openapiPath = m.input.Value()
				m.step = stepDone

			case stepDone:
				selected := getSelectedInfra(m.infraList.Items())
				*m.projectConfig = config.ProjectConfig{
					ProjectName: m.project,
					ModuleName:  m.module,
					OpenAPIPath: m.openapiPath,
					UsePostgres: contains(selected, "Postgres"),
					UseRedis:    contains(selected, "Redis"),
				}

				return m, tea.Quit
			}
		case tea.KeySpace:
			if m.step == stepInfraChoice {
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

func getSelectedInfra(items []list.Item) []string {
	var result []string
	for _, it := range items {
		ii := it.(infraItem)
		if ii.used {
			result = append(result, ii.title)
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
