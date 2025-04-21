package app

import (
	"fmt"

	"github.com/apereiroc/go-memo/internal/debug"
	"github.com/charmbracelet/lipgloss"
)

// define enumerated views
type currentView uint8

const (
	groupView currentView = iota
	commandView
)

func renderGroupView(m *model) string {
	// groups
	s1 := headerStyle.Render("Groups") + "\n\n"

	for idx, group := range m.groups {
		line := group.name
		if index(idx) == m.selectedGroup {
			line = selectedStyle.Render("> " + line)
		}

		s1 += line + "\n"
	}

	s1 = groupBoxStyle.Render(s1)

	// preview
	s2 := headerStyle.Render("Preview") + "\n\n"

	g := m.groups[m.selectedGroup]
	for _, cmd := range g.cmds {
		s2 += commandPreviewStyle.Render(cmd.cmd) + "\n"
	}

	s2 = previewBoxStyle.Render(s2)

	s := lipgloss.JoinHorizontal(lipgloss.Top, s1, s2)

	return s
}

func renderCommandView(m *model) string {
	g := m.groups[m.selectedGroup]

	// descriptions
	s1 := headerStyle.Render("Description") + "\n\n"

	// commands
	s2 := headerStyle.Render("Browsing commands for ")
	s2 += selectedGroupStyle.Render(g.name) + "\n\n"
	// s2 += headerStyle.Render(". Choose a command:") + "\n\n"

	for idx, cmd := range g.cmds {
		line := cmd.cmd
		if index(idx) == m.selectedCmd {
			line = selectedStyle.Render("> " + line)
			s1 += descriptionStyle.Render(cmd.description) + "\n"
		}
		s2 += line + "\n"
	}

	s1 = descBoxStyle.Render(s1)
	s2 = commandBoxStyle.Render(s2)

	s := lipgloss.JoinHorizontal(lipgloss.Top, s1, s2)

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
