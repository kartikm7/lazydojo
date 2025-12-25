package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kartikm7/lazydojo/internals/models"
)

func main() {
	p := tea.NewProgram(models.InitRootModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		panic("Ghand lag gayi")
	}
}
