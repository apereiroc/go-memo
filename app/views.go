package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Follow strategy design pattern for scalability
// Implement view's behaviours under model's Update/View methods for each case
// Also need to implement how model.next()/prev() are handled (view-dependent behaviour)
type viewStrategy interface {
	update(m model, msg tea.Msg) (viewStrategy, model, tea.Cmd)
	view(m model) string
	// Handle next entry based on current view
	next(m model) model
	// Handle previous entry based on current view
	prev(m model) model
}

func (m model) View() string {
	// clean screen after exit
	if m.quitWithCmd {
		return ""
	}
	return m.view.view(m)
}
