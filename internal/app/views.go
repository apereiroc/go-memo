package app

import (
	"fmt"

	"github.com/apereiroc/go-memo/internal/debug"
)

// define enumerated views
type currentView uint8

const (
	groupView currentView = iota
	commandView
)

func renderGroupView(m *model) string {
	s := "choose a group\n\n"

	for idx, group := range m.groups {
		cursor := ""
		if index(idx) == m.selectedGroup {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, group.name)
	}

	return s
}

func renderCommandView(m *model) string {
	s := "choose a command\n\n"

	for idx, cmd := range m.groups[m.selectedGroup].cmds {
		cursor := ""
		if index(idx) == m.selectedCmd {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, cmd.cmd)
	}

	return s
}

func (m model) View() string {
	switch m.view {
	case commandView:
		return renderCommandView(&m)
	case groupView:
		return renderGroupView(&m)
	default:
		err := fmt.Errorf("unknown view: %d", m.view)
		debug.Error(err)
		panic(err)
	}
}
