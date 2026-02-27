package theme

import "github.com/charmbracelet/lipgloss"

const (
	CheckIcon      = "✓"
	CrossIcon      = "✗"
	InfoIcon       = "ℹ"
	WarningIcon    = "⚠"
	DebugIcon      = "[D]"
	ArrowRightIcon = "→"
	PlusIcon       = "+"
)

var (
	Background    = lipgloss.Color("#292A2B")
	Foreground    = lipgloss.Color("#CCCCCC")
	Cursor        = lipgloss.Color("#FFB86C")
	Black         = lipgloss.Color("#000000")
	Red           = lipgloss.Color("#B81E4A")
	Green         = lipgloss.Color("#13FFDC")
	Yellow        = lipgloss.Color("#FFB86C")
	Blue          = lipgloss.Color("#7DC1FF")
	Magenta       = lipgloss.Color("#B084EB")
	Cyan          = lipgloss.Color("#35FFDC")
	White         = lipgloss.Color("#FFFFFF")
	BrightBlack   = lipgloss.Color("#7A8181")
	BrightRed     = lipgloss.Color("#F76E6E")
	BrightGreen   = lipgloss.Color("#13FFDC")
	BrightYellow  = lipgloss.Color("#DAC26B")
	BrightBlue    = lipgloss.Color("#5CA7E4")
	BrightMagenta = lipgloss.Color("#FF9AC1")
	BrightCyan    = lipgloss.Color("#00C990")
	BrightWhite   = lipgloss.Color("#989FB1")
)

func GetAnsiStyle(color lipgloss.Color) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(color)
}

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
