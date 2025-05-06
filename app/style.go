package app

import "github.com/charmbracelet/lipgloss"

const (
	orange = "215"
	pink   = "205"
	blue   = "69"
)

const (
	boxHeight            = 8
	boxWidth             = 50
	boxPaddingVertical   = 1
	boxPaddingHorizontal = 2
	helpBoxWidth         = 102
)

var (
	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(pink)).
			Bold(true)

	selectedGroupStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(blue)).
				Bold(true)

	commandPreviewStyle = lipgloss.NewStyle().
				Bold(true)

	descriptionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(orange)).
				Italic(true)

	headerStyle = lipgloss.NewStyle().Bold(true)

	groupBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(boxPaddingVertical, boxPaddingHorizontal).
			Width(boxWidth).
			Height(boxHeight)

	previewBoxStyle = groupBoxStyle
	commandBoxStyle = groupBoxStyle
	descBoxStyle    = groupBoxStyle
	helpBoxStyle    = groupBoxStyle.UnsetHeight().
			Width(helpBoxWidth).
			UnsetPaddingBottom().
			UnsetPaddingTop()
)
