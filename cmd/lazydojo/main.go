package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor int
	items  []string
}

func initModel() model {
	return model{
		items:  []string{"dummy data", "just for", "mocking sakes"},
		cursor: 1,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "j":
			if m.cursor < len(m.items) {
				m.cursor++
			} else if m.cursor == len(m.items) {
				m.cursor = 1
			}
		case "k":
			if m.cursor > 1 {
				m.cursor--
			} else if m.cursor == 1 {
				m.cursor = len(m.items)
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	// header
	s := "What task do you plan on completing?\n\n"
	for idx, val := range m.items {
		cursor := " "
		if idx == m.cursor-1 {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, val)
	}

	// footer
	s += "\nPress Q to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		panic("Ghand lag gayi")
	}
}
