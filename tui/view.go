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
	case stepInfraChoice:
		return m.infraList.View()
	case stepOpenAPIPath:
		return fmt.Sprintf("Input OpenAPI YAML path:\n\n%s\n\n[Enter] to finish", m.input.View())
	case stepDone:
		selected := getSelectedInfra(m.infraList.Items())

		return fmt.Sprintf(
			"✅ Done:\nProject: %s\nModule: %s\nOpenAPI: %s\nInfra: %s\n\n%s",
			m.project,
			m.module,
			m.openapiPath,
			strings.Join(selected, ", "),
			"Press [Enter] to generate project or [Esc] to quit",
		)
	default:
		return "Something went wrong"
	}
}
