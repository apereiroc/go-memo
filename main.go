package main

import (
	"log"

	debug "github.com/apereiroc/memogo/internal/debug"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct{}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return "hello world!"
}

func main() {
	debug.Start()
	defer debug.Stop()

	// start program
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		debug.Error(err)
		log.Fatal(err)
	}
}
