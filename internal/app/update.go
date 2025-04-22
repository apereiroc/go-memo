package app

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
		return m, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			return m.view.prev(m), nil
		case key.Matches(msg, m.keys.Down):
			return m.view.next(m), nil
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}

	// update view and model
	newView, newModel, cmd := m.view.update(m, msg)
	// attach the new view to the new model
	newModel.view = newView
	return newModel, cmd
}
