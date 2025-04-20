package app

import "github.com/charmbracelet/bubbles/key"

// keyMap defines a set of keybindings. To work for help it must satisfy
// key.Map. It could also very easily be a map[string]key.Binding.
type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Esc   key.Binding
	Quit  key.Binding
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "move down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "go next"),
	),
	Esc: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "go prev"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}
