package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type tickMsg time.Time // Custom tick message type

// Update is the Bubble Tea "update" function that dictates how the model changes
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Handle key input: quit the app on "q", "esc", or "ctrl+c"
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}

	case tickMsg: // Handle tick messages (countdown timer)
		if m.count <= 0 {
			return m, tea.Quit
		}
		m.count--
		return m, tick() // Schedule the next tick
	}

	return m, nil
}

// tick creates a command to send a tick message every second
func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
