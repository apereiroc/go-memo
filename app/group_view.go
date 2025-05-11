package app

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Rendering machinery when the user is selecting a group of commands
type groupView struct{}

func (v groupView) update(m model, msg tea.Msg) (viewStrategy, model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Enter):
			// update view
			m.view = commandView{}
			// update keymap
			m.keys = commandKeys
			// reset selection index
			m.selectedCmd = 0

			return m.view, m, nil
		}
	}
	return v, m, nil
}

func (v groupView) view(m model) string {
	// groups
	s1 := headerStyle.Render("Groups") + "\n\n"

	for idx, group := range m.groups {
		line := group.Name
		if index(idx) == m.selectedGroup {
			line = selectedStyle.Render("> " + line)
		}

		s1 += line + "\n"
	}

	s1 = groupBoxStyle.Render(s1)

	// preview
	s2 := headerStyle.Render("Preview") + "\n\n"

	g := m.groups[m.selectedGroup]
	for _, cmd := range g.Cmds {
		s2 += commandPreviewStyle.Render(cmd.Cmd) + "\n"
	}

	s2 = previewBoxStyle.Render(s2)

	s := lipgloss.JoinHorizontal(lipgloss.Top, s1, s2) + "\n"

	// print help
	s += helpBoxStyle.Render(m.help.View(m.keys))

	return s
}

func (v groupView) next(m model) model {
	// we're viewing the groups
	// need to access to the maximum number of groups
	maxGroups := index(len(m.groups))
	// advance to the end, and go to the beginning if the length is exceeded
	m.selectedGroup = (m.selectedGroup + 1) % maxGroups

	return m
}

func (v groupView) prev(m model) model {
	// we're viewing the groups
	switch m.selectedGroup {
	case 0:
		// go to the end if the user selects one prior to 0
		maxGroups := index(len(m.groups))
		m.selectedGroup = maxGroups - 1
	default:
		m.selectedGroup--
	}

	return m
}
