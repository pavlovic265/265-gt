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

// Panda Terminal Theme Colors
// Based on https://github.com/PandaTheme/panda-terminal
var (
	// Base colors
	Background = lipgloss.Color("#292A2B") // Panda background
	Foreground = lipgloss.Color("#CCCCCC") // Panda foreground
	Cursor     = lipgloss.Color("#FFB86C") // Panda cursor (yellow)

	// Panda Theme colors
	Black   = lipgloss.Color("#000000") // black
	Red     = lipgloss.Color("#EC2864") // red
	Green   = lipgloss.Color("#13FFDC") // green
	Yellow  = lipgloss.Color("#FFB86C") // yellow
	Blue    = lipgloss.Color("#7DC1FF") // blue
	Magenta = lipgloss.Color("#B084EB") // magenta (purple)
	Cyan    = lipgloss.Color("#35FFDC") // cyan
	White   = lipgloss.Color("#FFFFFF") // white

	// Bright colors
	BrightBlack   = lipgloss.Color("#7A8181") // bright_black
	BrightRed     = lipgloss.Color("#F76E6E") // bright_red
	BrightGreen   = lipgloss.Color("#13FFDC") // bright_green
	BrightYellow  = lipgloss.Color("#DAC26B") // bright_yellow
	BrightBlue    = lipgloss.Color("#5CA7E4") // bright_blue
	BrightMagenta = lipgloss.Color("#FF9AC1") // bright_magenta (pink)
	BrightCyan    = lipgloss.Color("#00C990") // bright_cyan
	BrightWhite   = lipgloss.Color("#989FB1") // bright_white
)

// Helper functions to get Panda colors
func GetAnsiColor(ansiNumber int) lipgloss.Color {
	switch ansiNumber {
	case 0:
		return Black
	case 1:
		return Red
	case 2:
		return Green
	case 3:
		return Yellow
	case 4:
		return Blue
	case 5:
		return Magenta
	case 6:
		return Cyan
	case 7:
		return White
	case 8:
		return BrightBlack
	case 9:
		return BrightRed
	case 10:
		return BrightGreen
	case 11:
		return BrightYellow
	case 12:
		return BrightBlue
	case 13:
		return BrightMagenta
	case 14:
		return BrightCyan
	case 15:
		return BrightWhite
	default:
		return Foreground
	}
}

// Convenience functions for Panda colors
func GetBackgroundColor() lipgloss.Color {
	return Background
}

func GetForegroundColor() lipgloss.Color {
	return Foreground
}

func GetCursorColor() lipgloss.Color {
	return Cursor
}

// ANSI color styles for common use cases
func GetAnsiStyle(ansiNumber int) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(GetAnsiColor(ansiNumber))
}

// Common ANSI color combinations
func GetSuccessAnsiStyle() lipgloss.Style {
	return GetAnsiStyle(2) // Green
}

func GetErrorAnsiStyle() lipgloss.Style {
	return GetAnsiStyle(1) // Red
}

func GetWarningAnsiStyle() lipgloss.Style {
	return GetAnsiStyle(3) // Yellow
}

func GetInfoAnsiStyle() lipgloss.Style {
	return GetAnsiStyle(4) // Blue
}

func GetDebugAnsiStyle() lipgloss.Style {
	return GetAnsiStyle(5) // Magenta
}
