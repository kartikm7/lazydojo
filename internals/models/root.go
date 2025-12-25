// Package models is used to bifurcate each UI model, into it's own separate file
// models include Root and Home at the moment
package models

import tea "github.com/charmbracelet/bubbletea"

// rootModel, is just a wrapper to allow switching between models
type rootModel struct {
	model tea.Model
}

func InitRootModel() rootModel {
	initScreen := InitHomeModel()
	return rootModel{model: initScreen}
}

func (m rootModel) Init() tea.Cmd {
	return m.model.Init()
}

func (m rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.model.Update(msg)
}

func (m rootModel) View() string {
	// TODO: Add help footer here
	return m.model.View()
}
