package models

// basic element of the application
// stores a command and a description of what it does
type Command struct {
	Cmd         string // actual command
	Description string // description about what the command does
}
