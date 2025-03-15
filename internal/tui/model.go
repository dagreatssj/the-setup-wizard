package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Model represents the state of the TUI
type model struct {
	count int
}

// InitialModel initializes the application model
func initialModel() model {
	return model{count: 5} // Start counting down from 5
}

// Init sets up any initial commands for the program
func (m model) Init() tea.Cmd {
	return tick() // Start the timer
}
