package tui

import (
	"fmt"
	"strings"
)

func (m *Model) View() string {
	switch m.step {
	case stepProjectName:
		return fmt.Sprintf("Input the project name:\n\n%s\n\n[Enter] to continue", m.input.View())
	case stepModuleName:
		return fmt.Sprintf("Input Go-module name:\n\n%s\n\n[Enter] to continue", m.input.View())
	case stepOpenAPIPath:
		return fmt.Sprintf("Input OpenAPI YAML path:\n\n%s\n\n[Enter] to finish", m.input.View())
	case stepDBChoice:
		return m.dbList.View()
	case stepInfraChoice:
		return m.infraList.View()
	case stepDone:
		selected := getSelectedInfra(m.infraList.Items())
		db := getSelectedDB(m.dbList.Items())
		return fmt.Sprintf(
			"Project: %s\nModule: %s\nOpenAPI: %s\nDB: %s\nInfra: %s\n\n%s",
			m.project,
			m.module,
			m.openapiPath,
			db,
			strings.Join(selected, ", "),
			"Press [Enter] to generate project or [Esc] to quit",
		)
	default:
		return "Something went wrong"
	}
}
