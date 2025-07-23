package config

import "github.com/charmbracelet/lipgloss"

// Color palette based on purpose - shared across all commands
var (
	// Success - Turquoise-Green
	SuccessStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#19F9D8")).
		Bold(true)

	// Error - Pink
	ErrorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF76B5")).
		Bold(true)

	// Warning - Orange
	WarningStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F7B36A")).
		Bold(true)

	// Info - Light Blue
	InfoStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6FC1FF")).
		Bold(true)

	// Debug/Meta - Light Purple
	DebugStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#B180D7")).
		Bold(true)

	// Highlight - Light Pink
	HighlightStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF87B4")).
		Bold(true)

	// Title with highlight background
	TitleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#19F9D8")). // Success color for title
		Background(lipgloss.Color("#FF87B4")). // Highlight background
		Padding(0, 1).
		Bold(true)
) 