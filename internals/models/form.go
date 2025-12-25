package models

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type formModel struct {
	textInput textinput.Model
}

func InitFormModel() formModel {
	ti := textinput.New()
	ti.Placeholder = "What you practicing at the Dojo?"
	ti.Focus()
	ti.Width = 50
	return formModel{ti}
}

func (m formModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m formModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	updated, globalcmd := DefaultBinding(msg, m)
	if assert, ok := updated.(formModel); ok {
		m = assert
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, tea.Batch(globalcmd, cmd)
}

func (m formModel) View() string {
	return fmt.Sprintf("Form\n%s", m.textInput.View())
}
