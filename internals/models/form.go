package models

import (
	"database/sql"
	"fmt"

	catppuccin "github.com/catppuccin/go"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kartikm7/lazydojo/internals/models/helpbars"
	"github.com/kartikm7/lazydojo/pkg/db"
)

type formModel struct {
	keys      helpbars.AddTaskKeyMap
	help      help.Model
	textInput textinput.Model
	db        *sql.DB
}

func InitFormModel(db *sql.DB) formModel {
	width, _ := GetTermSize()
	ti := textinput.New()
	ti.Placeholder = "What you practicing at the Dojo?"
	ti.Focus()
	ti.Width = width / 2
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(catppuccin.Latte.Text().Hex))
	return formModel{textInput: ti, db: db, keys: helpbars.AddTaskKeys, help: helpbars.CreateAddTaskHelpBar()}
}

func (m formModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m formModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	updated, globalcmd := DefaultBinding(msg, m, m.db)
	if _, ok := updated.(formModel); !ok && !m.textInput.Focused() {
		return updated, globalcmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			if m.textInput.Focused() {
				globalcmd = nil
			}
		}
		switch msg.Type {
		case tea.KeyEnter:
			db.Add(m.db, m.textInput.Value())
			return InitHomeModel(m.db), nil
		case tea.KeyEsc:
			m.textInput.Blur()
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, tea.Batch(globalcmd, cmd)
}

func (m formModel) View() string {
	width, height := GetTermSize()
	help := lipgloss.NewStyle().Width(width)
	helpHeight := help.GetHeight()
	helpView := m.help.View(m.keys)
	parent := lipgloss.NewStyle().Width(width).Height(height - helpHeight).AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center)
	text := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color(catppuccin.Latte.Lavender().Hex)).Padding(1, 2)
	return fmt.Sprint(parent.Render(text.Render(m.textInput.View())) + "\n" + help.Render(helpView))
}
