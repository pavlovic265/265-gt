package log

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/ui/theme"
)

var (
	successIconStyle = lipgloss.NewStyle().
				Foreground(theme.Green)

	errorIconStyle = lipgloss.NewStyle().
			Foreground(theme.Red)

	infoIconStyle = lipgloss.NewStyle().
			Foreground(theme.Blue)

	warningIconStyle = lipgloss.NewStyle().
				Foreground(theme.Yellow)

	messageStyle = lipgloss.NewStyle().
			Foreground(theme.White)
)

func Info(message string) {
	fmt.Printf("%s %s\n",
		infoIconStyle.Render(theme.InfoIcon),
		messageStyle.Render(message))
}

func Error(message string, err error) error {
	return fmt.Errorf("%s %s: %w",
		errorIconStyle.Render(theme.CrossIcon),
		messageStyle.Render(message),
		err)
}

func ErrorMsg(message string) error {
	return fmt.Errorf("%s %s",
		errorIconStyle.Render(theme.CrossIcon),
		messageStyle.Render(message))
}

func Success(message string) {
	fmt.Printf("%s %s\n",
		successIconStyle.Render(theme.CheckIcon),
		messageStyle.Render(message))
}

func Warning(message string) {
	fmt.Printf("%s %s\n",
		warningIconStyle.Render(theme.WarningIcon),
		messageStyle.Render(message))
}

func Errorf(format string, args ...any) error {
	message := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n",
		errorIconStyle.Render(theme.CrossIcon),
		messageStyle.Render(message))
	return fmt.Errorf("%s", message)
}

func Infof(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	Info(message)
}

func Successf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	Success(message)
}

func Warningf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	Warning(message)
}
