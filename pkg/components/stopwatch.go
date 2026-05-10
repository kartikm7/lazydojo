package components

import (
	"log/slog"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type stopwatchModel struct {
	elapsed   time.Duration
	startTime time.Time
	running   bool
}

type (
	TickMsg struct{}
	StopMsg struct{}
)

func New(spent int) stopwatchModel {
	return stopwatchModel{elapsed: time.Duration(spent)}
}

func (m stopwatchModel) Init() tea.Cmd {
	return nil
}

func (m stopwatchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.elapsed += time.Since(m.startTime)
	switch msg := msg.(type) {
	case TickMsg:
		cmd := tea.Tick(time.Second, func(time.Time) tea.Msg {
			slog.Info("I hope this is working")
			return TickMsg{}
		})
		return m, cmd
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeySpace:
			return m, m.Toggle()
		}
	}
	return m, nil
}

func (m stopwatchModel) View() string {
	return ""
}

func (m stopwatchModel) Start() tea.Cmd {
	// elapsed calculation
	m.running = true
	m.startTime = time.Now()
	// this get's the UI update going
	v := tea.Tick(time.Second, func(time.Time) tea.Msg {
		return TickMsg{}
	})
	return v
}

func (m stopwatchModel) Stop() tea.Cmd {
	m.running = false
	return func() tea.Msg {
		return StopMsg{}
	}
}

func (m stopwatchModel) Toggle() tea.Cmd {
	if m.running {
		return m.Stop()
	} else {
		return m.Start()
	}
}
