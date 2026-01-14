package log

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/constants"
)

var (
	successIconStyle = lipgloss.NewStyle().
				Foreground(constants.Green)

	errorIconStyle = lipgloss.NewStyle().
			Foreground(constants.Red)

	infoIconStyle = lipgloss.NewStyle().
			Foreground(constants.Blue)

	warningIconStyle = lipgloss.NewStyle().
				Foreground(constants.Yellow)

	messageStyle = lipgloss.NewStyle().
			Foreground(constants.White)
)

func Info(message string) {
	fmt.Printf("%s %s\n",
		infoIconStyle.Render(constants.InfoIcon),
		messageStyle.Render(message))
}

func Error(message string, err error) error {
	return fmt.Errorf("%s %s: %w",
		errorIconStyle.Render(constants.CrossIcon),
		messageStyle.Render(message),
		err)
}

func ErrorMsg(message string) error {
	return fmt.Errorf("%s %s",
		errorIconStyle.Render(constants.CrossIcon),
		messageStyle.Render(message))
}

func Success(message string) {
	fmt.Printf("%s %s\n",
		successIconStyle.Render(constants.CheckIcon),
		messageStyle.Render(message))
}

func Warning(message string) {
	fmt.Printf("%s %s\n",
		warningIconStyle.Render(constants.WarningIcon),
		messageStyle.Render(message))
}

func Errorf(format string, args ...any) error {
	message := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n",
		errorIconStyle.Render(constants.CrossIcon),
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
