package models

import (
	"database/sql"
	"github.com/charmbracelet/log"
	"strconv"

	catppuccingo "github.com/catppuccin/go"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	query "github.com/kartikm7/lazydojo/pkg/db"
)

// keyMap defines a set of keybindings. To work for help it must satisfy
// key.Map. It could also very easily be a map[string]key.Binding.
type keyMap struct {
	Up      key.Binding
	Down    key.Binding
	Left    key.Binding
	Right   key.Binding
	Help    key.Binding
	Quit    key.Binding
	Home    key.Binding
	Timer   key.Binding
	AddTask key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Home, k.Timer, k.AddTask, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right},              // first column
		{k.Help, k.Home, k.Timer, k.AddTask, k.Quit}, // second column
	}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	), Left: key.NewBinding(key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Home: key.NewBinding(
		key.WithKeys("1"),
		key.WithHelp("1", "home"),
	),
	Timer: key.NewBinding(
		key.WithKeys("2"),
		key.WithHelp("2", "timer"),
	),
	AddTask: key.NewBinding(
		key.WithKeys("3"),
		key.WithHelp("3", "add task"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

type homeModel struct {
	keys       keyMap
	help       help.Model
	inputStyle lipgloss.Style
	db         *sql.DB
	table      table.Model
}

func InitHomeModel(db *sql.DB) homeModel {
	columns := []table.Column{{Title: "ID", Width: 4}, {Title: "Task", Width: 100}}
	tasks, err := query.ListEverything(db)
	if err != nil {
		// TODO: Add a notifier system
		log.Fatalf("Something went wrong dawg: %s", err)
	}

	// now we will covert tasks, by flattening it out to with the table.rows format
	rows := []table.Row{}
	for _, val := range tasks {
		// could've made this a one liner, but I think readability goes a longer way
		task := []string{strconv.Itoa(val.ID), val.Task}
		rows = append(rows, task)
	}

	table := table.New(table.WithColumns(columns), table.WithRows(rows), table.WithFocused(true))
	table.SetStyles(DefaultTableStyles())
	return homeModel{
		keys:       keys,
		help:       help.New(),
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
		default:
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
