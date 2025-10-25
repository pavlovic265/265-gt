package components

import (
	"fmt"

	"github.com/pavlovic265/265-gt/constants"
)

// ButtonType defines the visual style of a button
type ButtonType int

const (
	ButtonPrimary ButtonType = iota
	ButtonSuccess
	ButtonInfo
	ButtonWarning
	ButtonDanger
)

// Button represents a styled button component
type Button struct {
	Label      string
	Icon       string
	ButtonType ButtonType
	Focused    bool
}

// NewButton creates a new button with the given label and type
func NewButton(label string, buttonType ButtonType) Button {
	return Button{
		Label:      label,
		ButtonType: buttonType,
		Focused:    false,
	}
}

// WithIcon adds an icon to the button
func (b Button) WithIcon(icon string) Button {
	b.Icon = icon
	return b
}

// WithFocus sets the focus state of the button
func (b Button) WithFocus(focused bool) Button {
	b.Focused = focused
	return b
}

// Render returns the styled button string
func (b Button) Render() string {
	var text string
	if b.Icon != "" {
		text = fmt.Sprintf("[ %s %s ]", b.Icon, b.Label)
	} else {
		text = fmt.Sprintf("[ %s ]", b.Label)
	}

	// If not focused, return plain text
	if !b.Focused {
		return text
	}

	// Apply style based on button type when focused
	switch b.ButtonType {
	case ButtonSuccess:
		return constants.GetSuccessAnsiStyle().Render(text)
	case ButtonInfo:
		return constants.GetInfoAnsiStyle().Render(text)
	case ButtonWarning:
		return constants.GetWarningAnsiStyle().Render(text)
	case ButtonDanger:
		return constants.GetErrorAnsiStyle().Render(text)
	case ButtonPrimary:
		fallthrough
	default:
		return constants.GetAnsiStyle(constants.Blue).Bold(true).Render(text)
	}
}

// Common button constructors for convenience

// NewDoneButton creates a styled "Done" button with checkmark icon
func NewDoneButton(focused bool) Button {
	return NewButton("Done", ButtonSuccess).
		WithIcon(constants.CheckIcon).
		WithFocus(focused)
}

// NewAddButton creates a styled "Add" button with plus icon
func NewAddButton(focused bool) Button {
	return NewButton("Add", ButtonInfo).
		WithIcon(constants.PlusIcon).
		WithFocus(focused)
}

// NewCancelButton creates a styled "Cancel" button
func NewCancelButton(focused bool) Button {
	return NewButton("Cancel", ButtonDanger).
		WithFocus(focused)
}

// NewSaveButton creates a styled "Save" button
func NewSaveButton(focused bool) Button {
	return NewButton("Save", ButtonSuccess).
		WithFocus(focused)
}

// NewBackButton creates a styled "Back" button
func NewBackButton(focused bool) Button {
	return NewButton("Back", ButtonPrimary).
		WithFocus(focused)
}
