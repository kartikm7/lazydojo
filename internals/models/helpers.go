package models

import (
	"database/sql"

	tea "github.com/charmbracelet/bubbletea"
)

func DefaultBinding(msg tea.Msg, m tea.Model, db *sql.DB) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "1":
			return InitHomeModel(db), nil
		case "2":
			return InitFormModel(db), nil
		default:
			return m, nil
		}
	}
	return m, nil
}
