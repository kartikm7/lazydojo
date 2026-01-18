package models

import (
	"github.com/charmbracelet/log"
	"os"

	catppuccingo "github.com/catppuccin/go"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

func GetTermSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatalf("Something went wrong while finding size: %s", err)
	}
	return width, height
}

func DefaultTableStyles() table.Styles {
	style := table.DefaultStyles()
	style.Header = style.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(catppuccingo.Latte.Subtext1().Hex)).
		BorderBottom(true).
		Bold(false)

	style.Selected = style.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color(catppuccingo.Latte.Lavender().Hex)).
		Bold(false)
	return style
}
