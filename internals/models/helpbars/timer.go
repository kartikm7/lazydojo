package helpbars

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

// StopWatchKeyMap defines a set of keybindings. To work for help it must satisfy
// key.Map. It could also very easily be a map[string]key.Binding.
type StopWatchKeyMap struct {
	ToggleWatch key.Binding
	Help        key.Binding
	Quit        key.Binding
	Home        key.Binding
	Timer       key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k StopWatchKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.ToggleWatch, k.Help, k.Home, k.Timer, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k StopWatchKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.ToggleWatch},                   // first column
		{k.Help, k.Home, k.Timer, k.Quit}, // second column
	}
}

var StopWatchKeys = StopWatchKeyMap{
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
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

func CreateStopWatchHelpBar() help.Model {
	return help.New()
}
