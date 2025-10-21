package constants

import "github.com/charmbracelet/lipgloss"

// Status indicators with ASCII icons
const (
	CheckIcon      = "✓"
	CrossIcon      = "✗"
	InfoIcon       = "ℹ"
	WarningIcon    = "⚠"
	DebugIcon      = "[D]"
	ArrowRightIcon = "→"
	PlusIcon       = "+"
)

// Panda Terminal Theme Colors
// Based on https://github.com/PandaTheme/panda-terminal
var (
	// Base colors
	Background = lipgloss.Color("#292A2B") // Panda background
	Foreground = lipgloss.Color("#CCCCCC") // Panda foreground
	Cursor     = lipgloss.Color("#FFB86C") // Panda cursor (yellow)

	// Panda Theme colors
	Black   = lipgloss.Color("#000000") // black
	Red     = lipgloss.Color("#B81E4A") // red (darker)
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

// ANSI color styles for common use cases
func GetAnsiStyle(color lipgloss.Color) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(color)
}

// Common ANSI color combinations
func GetSuccessAnsiStyle() lipgloss.Style {
	return GetAnsiStyle(Green)
}

func GetErrorAnsiStyle() lipgloss.Style {
	return GetAnsiStyle(Red)
}

func GetWarningAnsiStyle() lipgloss.Style {
	return GetAnsiStyle(Yellow)
}

func GetInfoAnsiStyle() lipgloss.Style {
	return GetAnsiStyle(Blue)
}

func GetDebugAnsiStyle() lipgloss.Style {
	return GetAnsiStyle(Magenta)
}
