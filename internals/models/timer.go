package models

import (
	"database/sql"
	"fmt"

	catppuccin "github.com/catppuccin/go"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/stopwatch"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kartikm7/lazydojo/internals/models/helpbars"
)

type timerModel struct {
	taskRow   *table.Row
	keys      helpbars.StopWatchKeyMap
	help      help.Model
	stopwatch stopwatch.Model
	db        *sql.DB
}

func InitTimerModel(task *table.Row, db *sql.DB) timerModel {
	stopwatchModel := stopwatch.New()
	return timerModel{task, helpbars.StopWatchKeys, helpbars.CreateStopWatchHelpBar(), stopwatchModel, db}
}

func (m timerModel) Init() tea.Cmd {
	return nil
}

func (m timerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	updated, globalcmd := DefaultBinding(msg, m, m.db)
	if _, ok := updated.(timerModel); !ok {
		return updated, globalcmd
	}

	var toggle tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeySpace:
			toggle = m.stopwatch.Toggle()
		}
	}

	m.stopwatch, cmd = m.stopwatch.Update(msg)
	return m, tea.Batch(globalcmd, cmd, toggle)
}

func (m timerModel) View() string {
	width, height := GetTermSize()
	help := lipgloss.NewStyle().Width(width)
	helpHeight := help.GetHeight()
	helpView := m.help.View(m.keys)
	parent := lipgloss.NewStyle().Width(width).Height(height - helpHeight).AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center)
	textParent := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color(catppuccin.Latte.Lavender().Hex)).Padding(2, 6)
	heading := lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color(catppuccin.Latte.Lavender().Hex))
	task := (*m.taskRow)
	return fmt.Sprint(parent.Render(textParent.Render(heading.Render(task[1])+" · "+m.stopwatch.View())) + "\n" + help.Render(helpView))
}
