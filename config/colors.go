package config

import "github.com/charmbracelet/lipgloss"

// Color palette based on Panda Syntax theme - superminimal, dark with subtle colors
var (
	// Success - Green (Panda theme green)
	SuccessStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A6E22E")).
			Bold(true)

	// Error - Red (Panda theme red)
	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F92672")).
			Bold(true)

	// Warning - Orange (Panda theme orange)
	WarningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FD971F")).
			Bold(true)

	// Info - Blue (Panda theme blue)
	InfoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#66D9EF")).
			Bold(true)

	// Debug/Meta - Purple (Panda theme purple)
	DebugStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AE81FF")).
			Bold(true)

	// Highlight - Yellow (Panda theme yellow)
	HighlightStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E6DB74")).
			Bold(true)

	// Title with highlight background
	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A6E22E")). // Success color for title
			Background(lipgloss.Color("#E6DB74")). // Highlight background
			Padding(0, 1).
			Bold(true)
)

// Status indicators with ASCII icons
const (
	SuccessIcon    = "✓"
	ErrorIcon      = "✗"
	ArrowRightIcon = "→"
)

// SuccessIndicator returns a styled success message with checkmark icon
func SuccessIndicator(message string) string {
	return SuccessStyle.Render(SuccessIcon + " " + message)
}

// ErrorIndicator returns a styled error message with X icon
func ErrorIndicator(message string) string {
	return ErrorStyle.Render(ErrorIcon + " " + message)
}

// SuccessIconOnly returns just the styled success icon
func SuccessIconOnly() string {
	return SuccessStyle.Render(SuccessIcon)
}

// ErrorIconOnly returns just the styled error icon
func ErrorIconOnly() string {
	return ErrorStyle.Render(ErrorIcon)
}
