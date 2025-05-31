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

	// Progress indicator
	progressBar := fmt.Sprintf("Step %d of 6: ", m.step+1)
	for i := 0; i < 6; i++ {
		if i < int(m.step) {
			progressBar += "● " // Completed step
		} else if i == int(m.step) {
			progressBar += "○ " // Current step
		} else {
			progressBar += "· " // Future step
		}
	}

	// Error message display
	errDisplay := ""
	if m.errMsg != "" {
		errDisplay = fmt.Sprintf("\n\n❌ Error: %s", m.errMsg)
	}

	switch m.step {
	case stepProjectName:
		return fmt.Sprintf("%s\nProject Configuration\n\nInput the project name:\n(e.g., my-awesome-project)\n\n%s\n\n📝 The project name will be used as the directory name.\n💡 Must start with a letter and contain only letters, numbers, hyphens, or underscores.%s%s",
			progressBar,
			m.input.View(),
			errDisplay,
			navHelp)
	case stepModuleName:
		return fmt.Sprintf("%s\nModule Configuration\n\nInput Go-module name:\n(e.g., github.com/user/myproject)\n\n%s%s%s",
			progressBar,
			m.input.View(),
			errDisplay,
			navHelp)
	case stepOpenAPIPath:
		if m.fileBrowser != nil {
			return fmt.Sprintf("%s\nAPI Specification\n\n%s%s%s",
				progressBar,
				m.fileBrowser.View(),
				errDisplay,
				navHelp)
		}
		// Fallback to text input if file browser is not available
		return fmt.Sprintf("%s\nAPI Specification\n\nInput OpenAPI YAML path:\n(e.g., example/server.yml)\n\n%s%s%s",
			progressBar,
			m.input.View(),
			errDisplay,
			navHelp)
	case stepDBChoice:
		return fmt.Sprintf("%s\nDatabase Selection\n\n%s%s",
			progressBar,
			m.dbList.View(),
			navHelp)
	case stepInfraChoice:
		return fmt.Sprintf("%s\nInfrastructure Selection\n\n%s%s",
			progressBar,
			m.infraList.View(),
			navHelp)
	case stepTesty:
		return fmt.Sprintf("%s\nTesting Framework\n\n%s%s",
			progressBar,
			m.testyList.View(),
			navHelp)
	case stepDone:
		selected := getSelectedInfraCodes(m.infraList.Items())
		db := getSelectedDB(m.dbList.Items())
		useTesty := "disabled"
		if getSelectedTesty(m.testyList.Items()) {
			useTesty = "enabled"
		}

		// Create a more visually appealing summary
		return fmt.Sprintf(
			"%s\n✨ Summary of Your Configuration ✨\n\n"+
			"📁 Project:  %s\n"+
			"📦 Module:   %s\n"+
			"📄 OpenAPI:  %s\n"+
			"🗄️  Database: %s\n"+
			"🔧 Infra:    %s\n"+
			"🧪 Testy:    %s%s",
			progressBar,
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
