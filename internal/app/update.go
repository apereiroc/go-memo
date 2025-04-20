package app

import (
	"fmt"

	"github.com/apereiroc/go-memo/internal/debug"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	debug.Info(fmt.Sprintf("got message: %+v", msg))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			// The user pressed up
			m.prev()
		case key.Matches(msg, m.keys.Down):
			// The user pressed up
			m.next()
		case key.Matches(msg, m.keys.Enter):
			// The user pressed enter
			switch m.view {
			case groupView:
				m.view = commandView
			}
		case key.Matches(msg, m.keys.Esc):
			// The user pressed esc
			switch m.view {
			case commandView:
				m.view = groupView
			}
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
}
