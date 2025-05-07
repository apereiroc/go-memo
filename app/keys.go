package app

import "github.com/charmbracelet/bubbles/key"

// keyMap defines a set of keybindings.
// To work for help it must satisfy
// key.Map. It could also very easily be a map[string]key.Binding.

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Esc   key.Binding
	Quit  key.Binding
}

// groupView keymap
var groupKeys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up", "shift+tab"),     // actual keybindings
		key.WithHelp("↑/k/shift+tab", "move up"), // corresponding help text
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down", "tab"),
		key.WithHelp("↓/j/tab", "move down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "go selected"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "quit"),
	),
}

// commandView keymap
var commandKeys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up", "shift+tab"),     // actual keybindings
		key.WithHelp("↑/k/shift+tab", "move up"), // corresponding help text
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down", "tab"),
		key.WithHelp("↓/j/tab", "move down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "copy and exit"),
	),
	Esc: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "go back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "quit"),
	),
}
