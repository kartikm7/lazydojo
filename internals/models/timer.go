package models

import (
	"database/sql"

	tea "github.com/charmbracelet/bubbletea"
)

type timerModel struct{ db *sql.DB }

func InitTimerModel(db *sql.DB) timerModel {
	return timerModel{db}
}

func (m timerModel) Init() tea.Cmd {
	return nil
}

func (m timerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return DefaultBinding(msg, m, m.db)
}

func (m timerModel) View() string {
	return "Timer screen"
}
