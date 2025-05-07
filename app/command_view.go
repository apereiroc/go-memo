package app

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Rendering machinery when the user is selecting a command within a group
type (
	commandView struct{}
)

func (v commandView) update(m model, msg tea.Msg) (viewStrategy, model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Esc):
			// update view
			m.view = groupView{}
			// update keymap
			m.keys = groupKeys

			return m.view, m, nil
		case key.Matches(msg, m.keys.Enter):
			// exit and copy the selected command to the clipboard
			m.quitWithCmd = true
			return v, m, tea.Quit
		}
	}
	return v, m, nil
}

func (v commandView) view(m model) string {
	// get select group
	g := m.groups[m.selectedGroup]

	// descriptions
	// s1 is the text that will be printed to descBoxStyle
	s1 := headerStyle.Render("Description") + "\n\n"

	// commands
	// s2 is the text that will be printed to commandBoxStyle
	s2 := headerStyle.Render("Browsing commands for ") +
		selectedGroupStyle.Render(g.Name) + "\n\n"

	for idx, cmd := range g.Cmds {
		line := cmd.Cmd
		if index(idx) == m.selectedCmd {
			// print description of the selected command
			s1 += descriptionStyle.Render(cmd.Description) + "\n"
			// highlight command if it matches the selected command
			line = selectedStyle.Render("> " + line)
		}
		// print command
		s2 += line + "\n"
	}

	s1 = descBoxStyle.Render(s1)
	s2 = commandBoxStyle.Render(s2)

	s := lipgloss.JoinHorizontal(lipgloss.Top, s1, s2) + "\n"

	// print help
	s += helpBoxStyle.Render(m.help.View(m.keys))

	return s
}

func (v commandView) next(m model) model {
	// we're viewing the commands
	// need to access the number of commands for the current group
	maxCmds := index(len(m.groups[m.selectedGroup].Cmds))
	// advance to the end, and go to the beginning if the length is exceeded
	m.selectedCmd = (m.selectedCmd + 1) % maxCmds

	return m
}

func (v commandView) prev(m model) model {
	// we're viewing the commands
	switch m.selectedCmd {
	case 0:
		// go to the end if the user selects one prior to 0
		maxCmds := index(len(m.groups[m.selectedGroup].Cmds))
		m.selectedCmd = maxCmds - 1
	default:
		m.selectedCmd--
	}
	return m
}
