package models

import tea "github.com/charmbracelet/bubbletea"

type formModel struct{}

func InitFormModel() formModel {
	return formModel{}
}

func (m formModel) Init() tea.Cmd {
	return nil
}

func (m formModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return DefaultBinding(msg, m)
}

func (m formModel) View() string {
	return "Form"
}
