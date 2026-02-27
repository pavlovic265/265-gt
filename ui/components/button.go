package components

import (
	"fmt"

	"github.com/pavlovic265/265-gt/ui/theme"
)

type ButtonType int

const (
	ButtonPrimary ButtonType = iota
	ButtonSuccess
	ButtonInfo
	ButtonWarning
	ButtonDanger
)

type Button struct {
	Label      string
	Icon       string
	ButtonType ButtonType
	Focused    bool
}

func NewButton(label string, buttonType ButtonType) Button {
	return Button{
		Label:      label,
		ButtonType: buttonType,
		Focused:    false,
	}
}

func (b Button) WithIcon(icon string) Button {
	b.Icon = icon
	return b
}

func (b Button) WithFocus(focused bool) Button {
	b.Focused = focused
	return b
}

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
		return theme.GetSuccessAnsiStyle().Render(text)
	case ButtonInfo:
		return theme.GetInfoAnsiStyle().Render(text)
	case ButtonWarning:
		return theme.GetWarningAnsiStyle().Render(text)
	case ButtonDanger:
		return theme.GetErrorAnsiStyle().Render(text)
	case ButtonPrimary:
		fallthrough
	default:
		return theme.GetAnsiStyle(theme.Blue).Bold(true).Render(text)
	}
}

func NewDoneButton(focused bool) Button {
	return NewButton("Done", ButtonSuccess).
		WithIcon(theme.CheckIcon).
		WithFocus(focused)
}

func NewAddButton(focused bool) Button {
	return NewButton("Add", ButtonInfo).
		WithIcon(theme.PlusIcon).
		WithFocus(focused)
}

func NewCancelButton(focused bool) Button {
	return NewButton("Cancel", ButtonDanger).
		WithFocus(focused)
}

func NewSaveButton(focused bool) Button {
	return NewButton("Save", ButtonSuccess).
		WithFocus(focused)
}

func NewBackButton(focused bool) Button {
	return NewButton("Back", ButtonPrimary).
		WithFocus(focused)
}

func NewSkipButton(focused bool) Button {
	return NewButton("Skip", ButtonWarning).
		WithFocus(focused)
}
