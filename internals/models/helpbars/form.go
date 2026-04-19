package helpbars

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

// AddTaskKeyMap defines a set of keybindings. To work for help it must satisfy
// key.Map. It could also very easily be a map[string]key.Binding.
type AddTaskKeyMap struct {
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
func (k AddTaskKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Home, k.Timer, k.AddTask, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k AddTaskKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right},              // first column
		{k.Help, k.Home, k.Timer, k.AddTask, k.Quit}, // second column
	}
}

var AddTaskKeys = AddTaskKeyMap{
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
	AddTask: key.NewBinding(
		key.WithKeys("2"),
		key.WithHelp("2", "add a task"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

func CreateAddTaskHelpBar() help.Model {
	return help.New()
}
