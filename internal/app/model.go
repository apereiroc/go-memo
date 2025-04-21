package app

import (
	"fmt"

	"github.com/apereiroc/go-memo/internal/debug"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

// basic element of the application
// stores a command and a description of what it does
type command struct {
	cmd         string // actual command
	description string // description about what the command does
}

// group of commands
// collection of commands under a common name
type group struct {
	name string    // name of the collection
	cmds []command // vector of commands
}

// Index to move through a group or command
type index uint32

// tea.Model is the main element of bubbletea
// its purpose is to hold the application's state
// three methods must be defined
// - Init() tea.Cmd
// - Update(tea.Msg) (tea.Model, tea.Cmd)
// - View() string
type model struct {
	groups        []group     // collection of group structures
	view          currentView // screen to be displayed in View()
	selectedGroup index       // index pointing to the current group
	selectedCmd   index       // index pointing to the current command
	keys          keyMap      // keys
	help          help.Model  // help
}

// required by bubbletea
// can return a Cmd that could perform some initial I/O.
// for now, we'll just return nil, which translates to "no command."
func (m model) Init() tea.Cmd {
	return nil
}

// model's creation
// this function will be called in main.go when creating the program
// it just needs to provide the initial state
func NewModel() model {
	m := model{
		groups: []group{
			{
				name: "C/C++",
				cmds: []command{
					{
						cmd:         "g++ --help",
						description: "Display the compiler's help message",
					},
					{
						cmd:         "g++ -o main main.cpp -I<include-path> -L<lib-path>",
						description: "Compile the source program, looking for possible headers in `include-path` and possible libraries in `lib-path`",
					},
				},
			},
			{
				name: "Git",
				cmds: []command{
					{
						cmd:         "git add <file>",
						description: "Add file contents to the index",
					},
					{
						cmd:         "git commit -m <message>",
						description: "Record changes to the repository",
					},
				},
			},
			{
				name: "Go",
				cmds: []command{
					{
						cmd:         "find . -type f -name '*.go' -exec sed -i '' 's/OLDNAME/NEWNAME/g' {} +",
						description: "Rename a go module from `OLDNAME` to `NEWNAME`",
					},
				},
			},
		},
		view:          groupView,
		selectedGroup: 0,
		selectedCmd:   0,
		keys:          groupKeys,
		help:          help.New(),
	}

	// customize help
	m.help.ShowAll = true        // show full help
	m.help.FullSeparator = " • " // add separator

	debug.Debug(fmt.Sprintf("initial model: %+v", m))
	return m
}

// Handle next entry based on current view
func (m *model) next() {
	switch m.view {
	case groupView:
		// we're viewing the groups
		// need to access to the maximum number of groups
		maxGroups := index(len(m.groups))
		// advance to the end, and go to the beginning if the length is exceeded
		m.selectedGroup = (m.selectedGroup + 1) % maxGroups
	case commandView:
		// we're viewing the commands
		// need to access the number of commands for the current group
		maxCmds := index(len(m.groups[m.selectedGroup].cmds))
		// advance to the end, and go to the beginning if the length is exceeded
		m.selectedCmd = (m.selectedCmd + 1) % maxCmds
	}
}

// Handle previous entry based on current view
func (m *model) prev() {
	switch m.view {
	case groupView:
		// we're viewing the groups
		switch m.selectedGroup {
		case 0:
			// go to the end if the user selects one prior to 0
			maxGroups := index(len(m.groups))
			m.selectedGroup = maxGroups - 1
		default:
			m.selectedGroup--
		}
	case commandView:
		// we're viewing the commands
		switch m.selectedCmd {
		case 0:
			// go to the end if the user selects one prior to 0
			maxCmds := index(len(m.groups[m.selectedGroup].cmds))
			m.selectedCmd = maxCmds - 1
		default:
			m.selectedCmd--
		}
	}
}
