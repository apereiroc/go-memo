package app

import (
	"fmt"
	"log"
	"os"

	"github.com/apereiroc/go-memo/debug"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	clipboard "github.com/tiagomelo/go-clipboard/clipboard"
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
	groups        []group      // collection of group structures
	view          viewStrategy // screen to be displayed in View()
	selectedGroup index        // index pointing to the current group
	selectedCmd   index        // index pointing to the current command
	keys          keyMap       // keys
	help          help.Model   // help
	quitWithCmd   bool         // whether the user closes the app (successfully selecting a command)
	quit          bool         // whether the user closes the app (abnormally)
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
			{
				name: "Docker",
				cmds: []command{
					{
						cmd:         "docker buildx build -f Dockerfile -t build .",
						description: "Build image from Dockerfile with tag build. Dockerfile is at .",
					},
					{
						cmd:         "docker buildx build --platform=linux/amd64 -f Dockerfile -t dotfiles-dev . && docker run --platform=linux/amd64 --rm -it dotfiles-dev",
						description: "Build image for testing my dotfiles. Hard to remember ...",
					},
				},
			},
		},
		view:          groupView{},
		selectedGroup: 0,
		selectedCmd:   0,
		keys:          groupKeys,
		help:          help.New(),
		quitWithCmd:   false,
		quit:          false,
	}

	// customize help
	m.help.ShowAll = true        // show full help
	m.help.FullSeparator = " • " // add separator

	debug.Debug(fmt.Sprintf("initial model: %+v", m))
	return m
}

// required by bubbletea
// can return a Cmd that could perform some initial I/O.
// for now, we'll just return nil, which translates to "no command."
func (m model) Init() tea.Cmd {
	return nil
}

// // Handle next entry based on current view
// func (m *model) next() {
// 	switch m.view {
// 	case groupView:
// 	case commandView:
// 	}
// }
//
// // Handle previous entry based on current view
// func (m *model) prev() {
// 	switch m.view {
// 	case groupView:
// 	case commandView:
// 	}
// }

// core function of the app
// if the model returned is correct and the quitting flag is set
// the selected command will be copied to the clipboard
func Success(progResult tea.Model) {
	// Type assertion here
	m, ok := progResult.(model)
	if !ok {
		log.Fatal("program result was not of type app.model")
	}

	if m.quitWithCmd {
		// successful exit
		// user selected a command
		c := clipboard.New()
		cmd := m.groups[m.selectedGroup].cmds[m.selectedCmd].cmd
		if err := c.CopyText(cmd); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
