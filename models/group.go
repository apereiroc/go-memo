package models

// group of commands
// collection of commands under a common name
type Group struct {
	Name string    // name of the collection
	Cmds []Command // vector of commands
}
