package main

import (
	"log"

	"github.com/apereiroc/go-memo/internal/app"
	"github.com/apereiroc/go-memo/internal/debug"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// comment these two lines for production
	// will early return and do nothing
	// debug.Start()
	// defer debug.Stop()

	// start program
	p := tea.NewProgram(app.InitialModel())
	if _, err := p.Run(); err != nil {
		debug.Error(err)
		log.Fatal(err)
	}
}
