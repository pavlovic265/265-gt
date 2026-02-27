package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/ui/theme"
)

type InputBuilder struct {
	model textinput.Model
}

func NewInput() *InputBuilder {
	m := textinput.New()

	// Set default styling
	m.Prompt = ""                                 // No prompt character
	m.Cursor.Style = theme.GetSuccessAnsiStyle()  // Green cursor
	m.PromptStyle = theme.GetSuccessAnsiStyle()   // Green prompt (for when prompt is set)
	m.TextStyle = theme.GetAnsiStyle(theme.White) // White text
	m.CharLimit = 256                             // Default character limit
	m.Width = 20                                  // Default width

	return &InputBuilder{model: m}
}

func (ib *InputBuilder) WithPlaceholder(placeholder string) *InputBuilder {
	ib.model.Placeholder = placeholder
	return ib
}

func (ib *InputBuilder) WithCharLimit(limit int) *InputBuilder {
	ib.model.CharLimit = limit
	return ib
}

func (ib *InputBuilder) WithWidth(width int) *InputBuilder {
	ib.model.Width = width
	return ib
}

func (ib *InputBuilder) WithFocus(focus bool) *InputBuilder {
	if focus {
		ib.model.Focus()
	} else {
		ib.model.Blur()
	}
	return ib
}

func (ib *InputBuilder) WithCursorStyle(style lipgloss.Color) *InputBuilder {
	ib.model.Cursor.Style = theme.GetAnsiStyle(style)
	return ib
}

func (ib *InputBuilder) WithPromptStyle(style lipgloss.Color) *InputBuilder {
	ib.model.PromptStyle = theme.GetAnsiStyle(style)
	return ib
}

func (ib *InputBuilder) WithTextStyle(style lipgloss.Color) *InputBuilder {
	ib.model.TextStyle = theme.GetAnsiStyle(style)
	return ib
}

func (ib *InputBuilder) Build() textinput.Model {
	return ib.model
}

func NewBranchInput() textinput.Model {
	return NewInput().
		WithPlaceholder("Branch").
		WithCharLimit(256).
		WithWidth(20).
		WithCursorStyle(theme.Yellow).
		Build()
}

func NewUserInput() textinput.Model {
	return NewInput().
		WithPlaceholder("User").
		WithCharLimit(32).
		WithWidth(50).
		WithCursorStyle(theme.Yellow).
		Build()
}

func NewTokenInput() textinput.Model {
	return NewInput().
		WithPlaceholder("Token").
		WithCharLimit(256).
		WithWidth(120).
		WithCursorStyle(theme.Yellow).
		Build()
}

func NewEmailInput() textinput.Model {
	return NewInput().
		WithPlaceholder("Email").
		WithCharLimit(100).
		WithWidth(60).
		WithCursorStyle(theme.Yellow).
		Build()
}

func NewNameInput() textinput.Model {
	return NewInput().
		WithPlaceholder("Full Name").
		WithCharLimit(100).
		WithWidth(60).
		WithCursorStyle(theme.Yellow).
		Build()
}

func NewSigningKeyInput() textinput.Model {
	return NewInput().
		WithPlaceholder("Signing Key (GPG)").
		WithCharLimit(100).
		WithWidth(60).
		WithCursorStyle(theme.Yellow).
		Build()
}

func NewGenericInput(placeholder string) textinput.Model {
	return NewInput().
		WithPlaceholder(placeholder).
		Build()
}
