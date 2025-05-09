package main

import (
	"log"

	"github.com/apereiroc/go-memo/app"
	"github.com/apereiroc/go-memo/debug"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// comment these two lines for production
	// will early return and do nothing
	// debug.Start()
	// defer debug.Stop()

	// start program
	p := tea.NewProgram(app.NewModel())
	if m, err := p.Run(); err != nil {
		debug.Error(err)
		log.Fatal(err)
	} else {
		app.Success(m)
	}
}
