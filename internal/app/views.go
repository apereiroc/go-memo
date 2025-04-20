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

func renderCommandView(m *model) string {
	return "hello from command view"
}

func renderGroupView(m *model) string {
	return "hello from group view"
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
