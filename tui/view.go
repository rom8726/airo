package tui

import (
	"fmt"
	"strings"
)

func (m *Model) View() string {
	// Common navigation help text
	var navHelp string
	if m.step > stepProjectName && m.step < stepDone {
		navHelp = "\n\n[Enter] continue • [Backspace] back • [Esc] quit"
	} else if m.step == stepProjectName {
		navHelp = "\n\n[Enter] continue • [Esc] quit"
	} else {
		navHelp = "\n\n[Enter] generate project • [Backspace] back • [Esc] quit"
	}

	// Progress indicator with step names
	stepNames := []string{"Project", "Module", "OpenAPI", "Database", "Infrastructure", "WebSocket+JWT", "Testing"}
	currentStep := int(m.step)

	var progressBar string
	if m.step == stepDone {
		progressBar = "Configuration complete!\n"
	} else {
		progressBar = fmt.Sprintf("Step %d of 7: %s\n", currentStep+1, stepNames[currentStep])
	}

	// Simplified progress bar
	progressBar += "["
	for i := 0; i < 7; i++ {
		if m.step == stepDone || i < currentStep {
			progressBar += "=" // Completed step
		} else if i == currentStep {
			progressBar += ">" // Current step
		} else {
			progressBar += "-" // Future step
		}
	}
	progressBar += "]"

	// Error message display
	errDisplay := ""
	if m.errMsg != "" {
		errDisplay = fmt.Sprintf("\n\nError: %s", m.errMsg)
	}

	switch m.step {
	case stepProjectName:
		return fmt.Sprintf("%s\nProject Configuration\n\nEnter project name (e.g., my-awesome-project):\n\n%s\n\nThe project name will be used as the directory name.\nMust start with a letter and contain only letters, numbers, hyphens, or underscores.%s%s",
			progressBar,
			m.input.View(),
			errDisplay,
			navHelp)
	case stepModuleName:
		return fmt.Sprintf("%s\nModule Configuration\n\nEnter Go module name (e.g., github.com/user/myproject):\n\n%s%s%s",
			progressBar,
			m.input.View(),
			errDisplay,
			navHelp)
	case stepOpenAPIPath:
		if !m.openapiDecisionMade {
			// Decision prompt: have a spec?
			return fmt.Sprintf("%s\nAPI Specification\n\nDo you already have an OpenAPI specification?\n\n[Y] Yes — choose a file\n[N] No — use the built-in ping-pong spec\n%s",
				progressBar,
				navHelp)
		}
		if m.fileBrowser != nil && !m.openapiUseEmbedded {
			return fmt.Sprintf("%s\nAPI Specification\n\nSelect your OpenAPI specification file:\n\n%s%s%s",
				progressBar,
				m.fileBrowser.View(),
				errDisplay,
				navHelp)
		}
		// Fallback to text input if file browser is not available
		if !m.openapiUseEmbedded {
			return fmt.Sprintf("%s\nAPI Specification\n\nEnter OpenAPI YAML path (e.g., example/server.yml):\n\n%s%s%s",
				progressBar,
				m.input.View(),
				errDisplay,
				navHelp)
		}
		return fmt.Sprintf("%s\nAPI Specification\n\nThe built-in ping-pong specification will be used.\n%s",
			progressBar,
			navHelp)
	case stepDBChoice:
		return fmt.Sprintf("%s\nDatabase Selection\n\nSelect a database for your project:\n\n%s%s",
			progressBar,
			m.dbList.View(),
			navHelp)
	case stepInfraChoice:
		return fmt.Sprintf("%s\nInfrastructure Selection\n\nSelect infrastructure components (optional):\n\n%s%s",
			progressBar,
			m.infraList.View(),
			navHelp)
	case stepRealtimeJWT:
		return fmt.Sprintf("%s\nWebSocket + JWT\n\nSelect WebSocket options:\n\n%s%s",
			progressBar,
			m.wsList.View(),
			navHelp)
	case stepTesty:
		return fmt.Sprintf("%s\nTesting Framework\n\nSelect testing options:\n\n%s%s",
			progressBar,
			m.testyList.View(),
			navHelp)
	case stepDone:
		selected := getSelectedInfraCodes(m.infraList.Items())
		db := getSelectedDB(m.dbList.Items())
		useRealtimeJWT := "No"
		if getSelectedRealtimeJWT(m.wsList.Items()) {
			useRealtimeJWT = "Yes"
		}
		useTesty := "No"
		if getSelectedTesty(m.testyList.Items()) {
			useTesty = "Yes"
		}

		// Truncate OpenAPI path if too long
		openAPIPath := m.openapiPath
		if len(openAPIPath) > 40 {
			openAPIPath = "..." + openAPIPath[len(openAPIPath)-37:]
		}

		// Format infra components
		infraStr := strings.Join(selected, ", ")
		if len(infraStr) == 0 {
			infraStr = "None"
		}
		if len(infraStr) > 40 {
			infraStr = infraStr[:37] + "..."
		}

		// Create a simple summary
		summary := "CONFIGURATION SUMMARY\n\n"
		summary += fmt.Sprintf("Project:      %s\n", m.project)
		summary += fmt.Sprintf("Module:       %s\n", m.module)
		summary += fmt.Sprintf("OpenAPI:      %s\n", openAPIPath)
		summary += fmt.Sprintf("Database:     %s\n", db)
		summary += fmt.Sprintf("Infra:        %s\n", infraStr)
		summary += fmt.Sprintf("WebSocket+JWT: %s\n", useRealtimeJWT)
		summary += fmt.Sprintf("Testy:        %s\n", useTesty)

		return fmt.Sprintf(
			"%s\n%s\n\nYour project is ready to be generated!%s",
			progressBar,
			summary,
			navHelp,
		)
	default:
		return "Something went wrong"
	}
}
