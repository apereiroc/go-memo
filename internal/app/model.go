package app

import tea "github.com/charmbracelet/bubbletea"

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

// tea.Model is the main element of bubbletea
// its purpose is to hold the application's state
// three methods must be defined
// - Init() tea.Cmd
// - Update(tea.Msg) (tea.Model, tea.Cmd)
// - View() string
type model struct {
	groups []group
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
func InitialModel() model {
	return model{
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
		},
	}
}
