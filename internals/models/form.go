package models

import (
	"database/sql"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kartikm7/lazydojo/pkg/db"
)

type formModel struct {
	textInput textinput.Model
	db        *sql.DB
}

func InitFormModel(db *sql.DB) formModel {
	ti := textinput.New()
	ti.Placeholder = "What you practicing at the Dojo?"
	ti.Focus()
	ti.Width = 50
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
			return m, nil
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
	return fmt.Sprintf("Form\n%s", m.textInput.View())
}
