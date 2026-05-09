package models

import (
	"database/sql"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/charmbracelet/log"

	catppuccingo "github.com/catppuccin/go"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kartikm7/lazydojo/internals/models/helpbars"
	database "github.com/kartikm7/lazydojo/pkg/db" // I don't like how it's not consistent might aswell name the package database
	"github.com/kartikm7/lazydojo/pkg/db/utils"
)

type homeModel struct {
	keys       helpbars.HomeKeyMap
	help       help.Model
	inputStyle lipgloss.Style
	db         *sql.DB
	table      table.Model
}

func InitHomeModel(db *sql.DB) homeModel {
	columns := []table.Column{{Title: "ID", Width: 4}, {Title: "Task", Width: 100}, {Title: "Completed", Width: 10}}
	tasks, err := database.ListEverything(db)
	if err != nil {
		// TODO: Add a notifier system
		log.Fatalf("Something went wrong dawg: %s", err)
	}

	// now we will covert tasks, by flattening it out to with the table.rows format
	rows := []table.Row{}
	for _, val := range tasks {
		// could've made this a one liner, but I think readability goes a longer way
		task := []string{strconv.Itoa(val.ID), val.Task, strconv.FormatBool(val.Completed)}
		rows = append(rows, task)
	}

	table := table.New(table.WithColumns(columns), table.WithRows(rows), table.WithFocused(true))
	table.SetStyles(DefaultTableStyles())
	return homeModel{
		keys:       helpbars.HomeKeys,
		help:       helpbars.CreateHomeHelpBar(),
		inputStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF75B7")),
		db:         db,
		table:      table,
	}
}

func (m homeModel) Init() tea.Cmd {
	return nil
}

func (m homeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// returning globalcmd
	updated, globalcmd := DefaultBinding(msg, m, m.db)
	if _, ok := updated.(homeModel); !ok {
		return updated, globalcmd
	}

	// checking for local values
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case m.keys.Help.Help().Key:
			m.help.ShowAll = !m.help.ShowAll
		case tea.KeyEnter.String():
			focusedValues := m.table.SelectedRow()
			return InitTimerModel(&focusedValues, m.db), nil
		case "x":
			focusedValues := m.table.SelectedRow()

			// getting the last stored value
			previousValue, err := strconv.ParseBool(focusedValues[2])
			if err != nil {
				slog.Error("Somehow the value was not a boolean, shouldn't be happening", "error", err)
			}
			focusedValues[2] = fmt.Sprintf("%v", !previousValue)
			m.table.UpdateViewport() // we update the just the table, THAT'S SO COOOl
			// now we run the DB updation
			idx, err := strconv.Atoi(focusedValues[0])
			if err != nil {
				slog.Info("Shit something went south", "error", err)
			}
			updatedValues := database.Values{}
			updatedValues.Completed = utils.Pointer(!previousValue)
			err = database.Update(m.db, idx, updatedValues)
			if err != nil {
				slog.Error("The update failed somehow the previous logs should help", "error", err)
			}
		default:
			slog.Info(msg.String())
		}
	}

	// this fetches the cmd for the table, and more importantly updates the table UI state
	m.table, cmd = m.table.Update(msg)

	return m, tea.Batch(globalcmd, cmd)
}

func (m homeModel) View() string {
	width, height := GetTermSize()
	parent := lipgloss.NewStyle().Width(width).Height(height).AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center)
	table := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color(catppuccingo.Latte.Subtext1().Hex))
	help := lipgloss.NewStyle().Width(width)
	helpView := m.help.View(m.keys)
	return parent.Render(table.Render(m.table.View())) + "\n" + help.Render(helpView)
}
