package main

import (
	"log"

	"github.com/apereiroc/go-memo/app"
	database "github.com/apereiroc/go-memo/db"
	"github.com/apereiroc/go-memo/debug"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// comment these two lines for production
	// will early return and do nothing
	// debug.Start()
	// defer debug.Stop()

	db, err := database.InitDB()
	if err != nil {
		debug.Error(err)
		log.Panic(err)
	}
	defer db.Close()

	// start program
	model, err := app.NewModel(db)
	if err != nil {
		debug.Error(err)
		log.Panic(err)
	}

	p := tea.NewProgram(model)
	if m, err := p.Run(); err != nil {
		debug.Error(err)
		log.Panic(err)
	} else {
		app.Success(m)
	}
}
