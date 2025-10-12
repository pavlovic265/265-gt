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

// 16-Color ANSI Palette (Dark Theme)
// Based on ~/.gtconfig.yaml theme configuration
var (
	// Dark theme colors
	DarkBackground = lipgloss.Color("#1d1e20") // distinct dark bg
	DarkForeground = lipgloss.Color("#e6e6e6") // distinct light fg
	DarkCursor     = lipgloss.Color("#ffb86c") // cursor

	// ANSI 0-7 (normal colors)
	DarkAnsi0 = lipgloss.Color("#1d1e20") // black
	DarkAnsi1 = lipgloss.Color("#ff5c57") // red
	DarkAnsi2 = lipgloss.Color("#A9DC52") // green
	DarkAnsi3 = lipgloss.Color("#f3f99d") // yellow
	DarkAnsi4 = lipgloss.Color("#57c7ff") // blue
	DarkAnsi5 = lipgloss.Color("#ff6ac1") // magenta
	DarkAnsi6 = lipgloss.Color("#9aedfe") // cyan
	DarkAnsi7 = lipgloss.Color("#e6e6e6") // white

	// ANSI 8-15 (bright colors)
	DarkAnsi8  = lipgloss.Color("#555555") // bright_black
	DarkAnsi9  = lipgloss.Color("#ff7a90") // bright_red
	DarkAnsi10 = lipgloss.Color("#69ff94") // bright_green
	DarkAnsi11 = lipgloss.Color("#ffffa5") // bright_yellow
	DarkAnsi12 = lipgloss.Color("#9aedfe") // bright_blue
	DarkAnsi13 = lipgloss.Color("#ff92d0") // bright_magenta
	DarkAnsi14 = lipgloss.Color("#c8ffff") // bright_cyan
	DarkAnsi15 = lipgloss.Color("#ffffff") // bright_white
)

// 16-Color ANSI Palette (Light Theme)
var (
	// Light theme colors
	LightBackground = lipgloss.Color("#fafafa")
	LightForeground = lipgloss.Color("#2e2e2e")
	LightCursor     = lipgloss.Color("#ff5c57")

	// ANSI 0-7 (normal colors)
	LightAnsi0 = lipgloss.Color("#2e2e2e") // black (text)
	LightAnsi1 = lipgloss.Color("#ff5c57") // red
	LightAnsi2 = lipgloss.Color("#3bb273") // green
	LightAnsi3 = lipgloss.Color("#d4b106") // yellow
	LightAnsi4 = lipgloss.Color("#268bd2") // blue
	LightAnsi5 = lipgloss.Color("#af5fff") // magenta
	LightAnsi6 = lipgloss.Color("#00bcd4") // cyan
	LightAnsi7 = lipgloss.Color("#fafafa") // white (bg-like)

	// ANSI 8-15 (bright colors)
	LightAnsi8  = lipgloss.Color("#999999") // bright_black
	LightAnsi9  = lipgloss.Color("#ff7b72") // bright_red
	LightAnsi10 = lipgloss.Color("#44d88d") // bright_green
	LightAnsi11 = lipgloss.Color("#ffe36e") // bright_yellow
	LightAnsi12 = lipgloss.Color("#4fc3f7") // bright_blue
	LightAnsi13 = lipgloss.Color("#c586c0") // bright_magenta
	LightAnsi14 = lipgloss.Color("#4dd0e1") // bright_cyan
	LightAnsi15 = lipgloss.Color("#ffffff") // bright_white
)

// Helper functions to get ANSI colors based on theme
func GetAnsiColor(ansiNumber int) lipgloss.Color {
	if isLightTheme() {
		return getLightAnsiColor(ansiNumber)
	}
	return getDarkAnsiColor(ansiNumber)
}

func getDarkAnsiColor(ansiNumber int) lipgloss.Color {
	switch ansiNumber {
	case 0:
		return DarkAnsi0
	case 1:
		return DarkAnsi1
	case 2:
		return DarkAnsi2
	case 3:
		return DarkAnsi3
	case 4:
		return DarkAnsi4
	case 5:
		return DarkAnsi5
	case 6:
		return DarkAnsi6
	case 7:
		return DarkAnsi7
	case 8:
		return DarkAnsi8
	case 9:
		return DarkAnsi9
	case 10:
		return DarkAnsi10
	case 11:
		return DarkAnsi11
	case 12:
		return DarkAnsi12
	case 13:
		return DarkAnsi13
	case 14:
		return DarkAnsi14
	case 15:
		return DarkAnsi15
	default:
		return DarkForeground
	}
}

func getLightAnsiColor(ansiNumber int) lipgloss.Color {
	switch ansiNumber {
	case 0:
		return LightAnsi0
	case 1:
		return LightAnsi1
	case 2:
		return LightAnsi2
	case 3:
		return LightAnsi3
	case 4:
		return LightAnsi4
	case 5:
		return LightAnsi5
	case 6:
		return LightAnsi6
	case 7:
		return LightAnsi7
	case 8:
		return LightAnsi8
	case 9:
		return LightAnsi9
	case 10:
		return LightAnsi10
	case 11:
		return LightAnsi11
	case 12:
		return LightAnsi12
	case 13:
		return LightAnsi13
	case 14:
		return LightAnsi14
	case 15:
		return LightAnsi15
	default:
		return LightForeground
	}
}

// Convenience functions for common ANSI colors
func GetBackgroundColor() lipgloss.Color {
	if isLightTheme() {
		return LightBackground
	}
	return DarkBackground
}

func GetForegroundColor() lipgloss.Color {
	if isLightTheme() {
		return LightForeground
	}
	return DarkForeground
}

func GetCursorColor() lipgloss.Color {
	if isLightTheme() {
		return LightCursor
	}
	return DarkCursor
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
