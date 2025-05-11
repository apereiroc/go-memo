package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	database "github.com/apereiroc/go-memo/db"
	"github.com/apereiroc/go-memo/debug"
	"github.com/apereiroc/go-memo/models"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	clipboard "github.com/tiagomelo/go-clipboard/clipboard"
)

// Index to move through a group or command
type index uint32

// tea.Model is the main element of bubbletea
// its purpose is to hold the application's state
// three methods must be defined
// - Init() tea.Cmd
// - Update(tea.Msg) (tea.Model, tea.Cmd)
// - View() string
type model struct {
	groups        []models.Group // collection of group structures
	view          viewStrategy   // screen to be displayed in View()
	selectedGroup index          // index pointing to the current group
	selectedCmd   index          // index pointing to the current command
	keys          keyMap         // keys
	help          help.Model     // help
	quitWithCmd   bool           // whether the user closes the app (successfully selecting a command)
	quit          bool           // whether the user closes the app (abnormally)
}

// model's creation
// this function will be called in main.go when creating the program
// it just needs to provide the initial state
func NewModel(db *sql.DB) (model, error) {
	groups, err := database.LoadGroups(db)
	if err != nil {
		return model{}, err
	}

	m := model{
		groups:        groups,
		selectedGroup: 0,
		selectedCmd:   0,
		keys:          groupKeys,
		help:          help.New(),
		quitWithCmd:   false,
		quit:          false,
	}

	// select initial view
	if len(groups) == 0 {
		m.view = noDatabaseView{}
	} else {
		m.view = groupView{}
	}

	// customize help
	m.help.ShowAll = true        // show full help
	m.help.FullSeparator = " • " // add separator

	debug.Debugf("initial model: %+v", m)
	return m, nil
}

// required by bubbletea
// can return a Cmd that could perform some initial I/O.
// for now, we'll just return nil, which translates to "no command."
func (m model) Init() tea.Cmd {
	return nil
}

// core function of the app
// if the model returned is correct and the quitting flag is set
// the selected command will be copied to the clipboard
func Success(progResult tea.Model) {
	// Type assertion here
	m, ok := progResult.(model)
	if !ok {
		log.Panic("program result was not of type app.model")
	}

	if m.quitWithCmd {
		// successful exit
		// user selected a command
		c := clipboard.New()
		cmd := m.groups[m.selectedGroup].Cmds[m.selectedCmd].Cmd
		if err := c.CopyText(cmd); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
