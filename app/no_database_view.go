package app

import tea "github.com/charmbracelet/bubbletea"

// Special screen to be displayed when the database is empty or not found
type noDatabaseView struct{}

func (v noDatabaseView) update(m *model, msg tea.Msg) (viewStrategy, *model, tea.Cmd) {
	return v, m, nil
}

func (v noDatabaseView) view(m *model) string {
	return emptyDatabase
}

func (v noDatabaseView) next(m *model) *model {
	return m
}

func (v noDatabaseView) prev(m *model) *model {
	return m
}
