package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func RunTeaApp() error {
	// Initialize program with the model
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	// Run the program and handle errors
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error running Bubble Tea app: %v", err)
		return err
	}
	return nil
}
