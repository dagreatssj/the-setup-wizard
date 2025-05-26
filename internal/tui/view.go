package tui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	var s string

	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1).
		Align(lipgloss.Center)

	if !m.dbSelected {
		s = viewDatabaseSelection(m)
	} else {
		s = viewScriptselection(m)
	}

	return style.Render(s)
}

func viewDatabaseSelection(m model) string {
	dbToSelect := m.dbChoice

	tpl := "Please select a database:\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("j/k, up/down: select") + dotStyle +
		subtleStyle.Render("enter: choose") + dotStyle +
		subtleStyle.Render("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s",
		checkbox("Local", dbToSelect == 0),
		checkbox("Google Firebase", dbToSelect == 1),
	)

	return fmt.Sprintf(tpl, choices)
}

func viewScriptselection(m model) string {
	return "script selection"
}
