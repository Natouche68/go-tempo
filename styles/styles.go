package styles

import (
	"github.com/charmbracelet/lipgloss"
	tsize "github.com/kopoli/go-terminal-size"
)

var Size, _ = tsize.GetSize()

var TitleStyle = lipgloss.NewStyle().
	Width(Size.Width).
	Align(lipgloss.Center).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("4")).
	BorderBottom(true).
	Bold(true).
	Foreground(lipgloss.Color("4"))

var ClockStyle = lipgloss.NewStyle().
	MarginTop(1).
	Bold(true).
	Foreground(lipgloss.Color("15"))

var LabelStyle = lipgloss.NewStyle().
	MarginBottom(2).
	Align(lipgloss.Center)

var WorkColor = lipgloss.Color("3")
var PauseColor = lipgloss.Color("2")

var HelpStyle = lipgloss.NewStyle().
	Width(Size.Width).
	Padding(0, 1).
	Align(lipgloss.Left).
	Italic(true).
	Foreground(lipgloss.Color("8"))

var ErrorStyle = lipgloss.NewStyle().
	Bold(true).
	Padding(0, 1).
	Background(lipgloss.Color("1"))
