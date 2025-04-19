package main

import (
	"log"
	"os"

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
	// log to file
	// useful because the TUI is blocking the terminal
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	// start program
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
