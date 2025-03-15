package tui

import "fmt"

// View is the Bubble Tea "view" function that renders the UI
func (m model) View() string {
	return fmt.Sprintf("\n\n     Hi. This program will exit in %d seconds...\n\nPress 'q' to quit.", m.count)
}
