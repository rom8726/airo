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

	// Progress indicator with step names
	stepNames := []string{"Project", "Module", "OpenAPI", "Database", "Infrastructure", "Testing"}
	currentStep := int(m.step)

	var progressBar string
	if m.step == stepDone {
		progressBar = "Configuration complete!\n"
	} else {
		progressBar = fmt.Sprintf("Step %d of 6: %s\n", currentStep+1, stepNames[currentStep])
	}
	progressBar += "["
	if m.step == stepDone {
		// All steps are completed
		for i := 0; i < 6; i++ {
			progressBar += "✅" // All steps completed
		}
	} else {
		for i := 0; i < 6; i++ {
			if i < currentStep {
				progressBar += "✅" // Completed step
			} else if i == currentStep {
				progressBar += "🔶" // Current step
			} else {
				progressBar += "⬜" // Future step
			}
		}
	}
	progressBar += "]"

	// Error message display
	errDisplay := ""
	if m.errMsg != "" {
		errDisplay = fmt.Sprintf("\n\n❌ ERROR ❌\n%s\n❌ ERROR ❌", m.errMsg)
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

		// Create a more visually appealing summary with a border
		border := "┌─────────────────────────────────────────────────────┐\n"
		border += "│                                                     │\n"
		border += "│            ✨ CONFIGURATION SUMMARY ✨              │\n"
		border += "│                                                     │\n"
		border += "├─────────────────────────────────────────────────────┤\n"

		// Format each line with proper spacing
		projectLine := fmt.Sprintf("│  📁 Project:  %-37s │\n", m.project)
		moduleLine := fmt.Sprintf("│  📦 Module:   %-37s │\n", m.module)

		// Truncate OpenAPI path if too long
		openAPIPath := m.openapiPath
		if len(openAPIPath) > 37 {
			openAPIPath = "..." + openAPIPath[len(openAPIPath)-34:]
		}
		openAPILine := fmt.Sprintf("│  📄 OpenAPI:  %-37s │\n", openAPIPath)

		dbLine := fmt.Sprintf("│  🗄️  Database: %-37s │\n", db)

		// Format infra components
		infraStr := strings.Join(selected, ", ")
		if len(infraStr) == 0 {
			infraStr = "none"
		}
		if len(infraStr) > 37 {
			infraStr = infraStr[:34] + "..."
		}
		infraLine := fmt.Sprintf("│  🔧 Infra:    %-37s │\n", infraStr)

		testyLine := fmt.Sprintf("│  🧪 Testy:    %-37s │\n", useTesty)

		border += projectLine
		border += moduleLine
		border += openAPILine
		border += dbLine
		border += infraLine
		border += testyLine
		border += "│                                                     │\n"
		border += "└─────────────────────────────────────────────────────┘\n"

		return fmt.Sprintf(
			"%s\n%s\n\n🎉 Your project is ready to be generated! 🎉%s",
			progressBar,
			border,
			navHelp,
		)
	default:
		return "Something went wrong"
	}
}
