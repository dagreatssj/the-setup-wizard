package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			return m, tea.Quit
		}
	}

	if !m.dbSelected {
		return updateDatabaseSelection(msg, m)
	}
	return updateScriptselection(msg, m)
}

func updateDatabaseSelection(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.dbChoice++
			if m.dbChoice > 1 {
				m.dbChoice = 1
			}
		case "k", "up":
			m.dbChoice--
			if m.dbChoice < 0 {
				m.dbChoice = 0
			}
		case "enter":
			m.dbSelected = true
			return m, nil
		}
	}

	return m, nil
}

func updateScriptselection(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	return m, nil
}
