package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

func DefaultBinding(msg tea.Msg, m tea.Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+1":
			return InitHomeModel(), nil
		case "ctrl+2":
			return InitTimerModel(), nil
		case "ctrl+3":
			return InitFormModel(), nil
		default:
			return m, nil
		}
	}
	return m, nil
}
