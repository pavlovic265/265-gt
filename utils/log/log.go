package log

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/constants"
)

// Styling definitions for log utility
var (
	// Success styles
	successIconStyle = lipgloss.NewStyle().
				Foreground(constants.Green)

	// Error styles
	errorIconStyle = lipgloss.NewStyle().
			Foreground(constants.Red)

	// Info styles
	infoIconStyle = lipgloss.NewStyle().
			Foreground(constants.Blue)

	// Warning styles
	warningIconStyle = lipgloss.NewStyle().
				Foreground(constants.Yellow)

	// Message styles
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

// Errorf creates a formatted error message and logs it
func Errorf(format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n",
		errorIconStyle.Render(constants.CrossIcon),
		messageStyle.Render(message))
	return fmt.Errorf("%s", message)
}

// Infof creates a formatted info message and logs it
func Infof(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	Info(message)
}

// Successf creates a formatted success message and logs it
func Successf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	Success(message)
}

// Warningf creates a formatted warning message and logs it
func Warningf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	Warning(message)
}
