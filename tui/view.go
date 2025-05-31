package tui

import (
	"fmt"
	"strings"
)

func (m *Model) View() string {
	// Common navigation help text
	var navHelp string
	if m.step > stepProjectName && m.step < stepDone {
		navHelp = "\n\n[Enter] to continue • [Backspace] to go back • [Esc] to quit"
	} else if m.step == stepProjectName {
		navHelp = "\n\n[Enter] to continue • [Esc] to quit"
	} else {
		navHelp = "\n\n[Enter] to generate project • [Backspace] to go back • [Esc] to quit"
	}

	// Error message display
	errDisplay := ""
	if m.errMsg != "" {
		errDisplay = fmt.Sprintf("\n\nError: %s", m.errMsg)
	}

	switch m.step {
	case stepProjectName:
		return fmt.Sprintf("Step 1 of 6: Project Configuration\n\nInput the project name:\n(e.g., my-awesome-project)\n\n%s%s%s",
			m.input.View(),
			errDisplay,
			navHelp)
	case stepModuleName:
		return fmt.Sprintf("Step 2 of 6: Module Configuration\n\nInput Go-module name:\n(e.g., github.com/user/myproject)\n\n%s%s%s",
			m.input.View(),
			errDisplay,
			navHelp)
	case stepOpenAPIPath:
		return fmt.Sprintf("Step 3 of 6: API Specification\n\nInput OpenAPI YAML path:\n(e.g., example/server.yml)\n\n%s%s%s",
			m.input.View(),
			errDisplay,
			navHelp)
	case stepDBChoice:
		return fmt.Sprintf("Step 4 of 6: Database Selection\n\n%s%s",
			m.dbList.View(),
			navHelp)
	case stepInfraChoice:
		return fmt.Sprintf("Step 5 of 6: Infrastructure Selection\n\n%s%s",
			m.infraList.View(),
			navHelp)
	case stepTesty:
		return fmt.Sprintf("Step 6 of 6: Testing Framework\n\n%s%s",
			m.testyList.View(),
			navHelp)
	case stepDone:
		selected := getSelectedInfraCodes(m.infraList.Items())
		db := getSelectedDB(m.dbList.Items())
		useTesty := ""
		if getSelectedTesty(m.testyList.Items()) {
			useTesty = "\nTesty framework: enabled"
		}

		return fmt.Sprintf(
			"Summary of Your Configuration:\n\nProject: %s\nModule: %s\nOpenAPI: %s\nDB: %s\nInfra: %s%s%s",
			m.project,
			m.module,
			m.openapiPath,
			db,
			strings.Join(selected, ", "),
			useTesty,
			navHelp,
		)
	default:
		return "Something went wrong"
	}
}
