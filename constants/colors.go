package constants

import "github.com/charmbracelet/lipgloss"

// Panda Syntax Dark Theme Colors (from https://github.com/PandaTheme/panda-syntax-vscode)
var (
	// Success - Green (Panda theme green)
	SuccessStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A9DC52")).
			Bold(true)

	// Error - Red (Panda theme red)
	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6188")).
			Bold(true)

	// Warning - Orange (Panda theme orange)
	WarningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD866")).
			Bold(true)

	// Info - Blue (Panda theme blue)
	InfoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#78DCE8")).
			Bold(true)

	// Debug/Meta - Purple (Panda theme purple)
	DebugStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AB9DF2")).
			Bold(true)

	// Highlight - Yellow (Panda theme yellow)
	HighlightStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD866")).
			Bold(true)

	// Command styling - Blue
	CommandStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#78DCE8")).
			Bold(true)

	// Branch styling - Purple
	BranchStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AB9DF2")).
			Bold(true)

	// File styling - Yellow italic
	FileStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD866")).
			Italic(true)

	// Status styling - Orange
	StatusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD866")).
			Bold(true)

	// Title with highlight background
	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A9DC52")). // Success color for title
			Background(lipgloss.Color("#FFD866")). // Highlight background
			Padding(0, 1).
			Bold(true)
)

// Panda Syntax Light Theme Colors (adapted for light terminals)
var (
	// Light Success - Dark Green
	LightSuccessStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#2D5016")).
				Bold(true)

	// Light Error - Dark Red
	LightErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8B1538")).
			Bold(true)

	// Light Warning - Dark Orange
	LightWarningStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#B8860B")).
				Bold(true)

	// Light Info - Dark Blue
	LightInfoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#1E3A8A")).
			Bold(true)

	// Light Debug - Dark Purple
	LightDebugStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6B46C1")).
			Bold(true)

	// Light Highlight - Dark Yellow
	LightHighlightStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#B8860B")).
				Bold(true)

	// Light Command styling - Dark Blue
	LightCommandStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#1E3A8A")).
				Bold(true)

	// Light Branch styling - Dark Purple
	LightBranchStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#6B46C1")).
				Bold(true)

	// Light File styling - Dark Yellow italic
	LightFileStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#B8860B")).
			Italic(true)

	// Light Status styling - Dark Orange
	LightStatusStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#B8860B")).
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

// Helper functions for themed styling based on config
func GetSuccessStyle() lipgloss.Style {
	if isLightTheme() {
		return LightSuccessStyle
	}
	return SuccessStyle
}

func GetErrorStyle() lipgloss.Style {
	if isLightTheme() {
		return LightErrorStyle
	}
	return ErrorStyle
}

func GetWarningStyle() lipgloss.Style {
	if isLightTheme() {
		return LightWarningStyle
	}
	return WarningStyle
}

func GetInfoStyle() lipgloss.Style {
	if isLightTheme() {
		return LightInfoStyle
	}
	return InfoStyle
}

func GetDebugStyle() lipgloss.Style {
	if isLightTheme() {
		return LightDebugStyle
	}
	return DebugStyle
}

func GetHighlightStyle() lipgloss.Style {
	if isLightTheme() {
		return LightHighlightStyle
	}
	return HighlightStyle
}

func GetCommandStyle() lipgloss.Style {
	if isLightTheme() {
		return LightCommandStyle
	}
	return CommandStyle
}

func GetBranchStyle() lipgloss.Style {
	if isLightTheme() {
		return LightBranchStyle
	}
	return BranchStyle
}

func GetFileStyle() lipgloss.Style {
	if isLightTheme() {
		return LightFileStyle
	}
	return FileStyle
}

func GetStatusStyle() lipgloss.Style {
	if isLightTheme() {
		return LightStatusStyle
	}
	return StatusStyle
}

// Theme detection based on config
func isLightTheme() bool {
	// For now, default to dark theme
	// TODO: This should be configurable through a proper config interface
	return false
}
