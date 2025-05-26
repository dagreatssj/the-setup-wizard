package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	dbChoice   int
	cursor     int
	dbSelected bool
}

func initialModel() model {
	return model{
		dbChoice:   0,
		cursor:     0,
		dbSelected: false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
