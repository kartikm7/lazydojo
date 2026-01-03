package models

import (
	"database/sql"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kartikm7/lazydojo/pkg/db"
)

type formModel struct {
	textInput textinput.Model
	db        *sql.DB
}

func InitFormModel(db *sql.DB) formModel {
	width, _ := GetTermSize()
	ti := textinput.New()
	ti.Placeholder = "What you practicing at the Dojo?"
	ti.Focus()
	ti.Width = width / 2
	return formModel{textInput: ti, db: db}
}

func (m formModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m formModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			db.Add(m.db, m.textInput.Value())
			return InitHomeModel(m.db), nil
		}
	}

	updated, globalcmd := DefaultBinding(msg, m, m.db)
	if assert, ok := updated.(formModel); ok {
		m = assert
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, tea.Batch(globalcmd, cmd)
}

func (m formModel) View() string {
	width, height := GetTermSize()
	parent := lipgloss.NewStyle().Width(width).Height(height).AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center)
	text := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#ea76cb")).Padding(1, 2)
	return fmt.Sprint(parent.Render(text.Render(m.textInput.View())))
}
