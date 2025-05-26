package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func RunWizardTeaApp() error {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error running Bubble Tea app: %v", err)
		return err
	}
	return nil
}
