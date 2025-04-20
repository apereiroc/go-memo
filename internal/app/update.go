package app

import (
	"fmt"

	"github.com/apereiroc/memogo/internal/debug"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	debug.Info(fmt.Sprintf("got message: %+v", msg))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}
