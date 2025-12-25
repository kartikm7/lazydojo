package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type timerModel struct{}

func InitTimerModel() timerModel {
	return timerModel{}
}

func (m timerModel) Init() tea.Cmd {
	return nil
}

func (m timerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return DefaultBinding(msg, m)
}

func (m timerModel) View() string {
	return "Timer screen"
}
