package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/constants"
)

// InputBuilder provides a fluent interface for creating styled text inputs
type InputBuilder struct {
	model textinput.Model
}

// NewInput creates a new input builder with sensible defaults
func NewInput() *InputBuilder {
	m := textinput.New()

	// Set default styling
	m.Prompt = ""                                         // No prompt character
	m.Cursor.Style = constants.GetSuccessAnsiStyle()      // Green cursor
	m.PromptStyle = constants.GetSuccessAnsiStyle()       // Green prompt (for when prompt is set)
	m.TextStyle = constants.GetAnsiStyle(constants.White) // White text
	m.CharLimit = 256                                     // Default character limit
	m.Width = 20                                          // Default width

	return &InputBuilder{model: m}
}

// WithPlaceholder sets the placeholder text
func (ib *InputBuilder) WithPlaceholder(placeholder string) *InputBuilder {
	ib.model.Placeholder = placeholder
	return ib
}

// WithCharLimit sets the character limit
func (ib *InputBuilder) WithCharLimit(limit int) *InputBuilder {
	ib.model.CharLimit = limit
	return ib
}

// WithWidth sets the input width
func (ib *InputBuilder) WithWidth(width int) *InputBuilder {
	ib.model.Width = width
	return ib
}

// WithFocus sets initial focus state
func (ib *InputBuilder) WithFocus(focus bool) *InputBuilder {
	if focus {
		ib.model.Focus()
	} else {
		ib.model.Blur()
	}
	return ib
}

// WithCursorStyle sets the cursor style color
func (ib *InputBuilder) WithCursorStyle(style lipgloss.Color) *InputBuilder {
	ib.model.Cursor.Style = constants.GetAnsiStyle(style)
	return ib
}

// WithPromptStyle sets the prompt style color
func (ib *InputBuilder) WithPromptStyle(style lipgloss.Color) *InputBuilder {
	ib.model.PromptStyle = constants.GetAnsiStyle(style)
	return ib
}

// WithTextStyle sets the text style color
func (ib *InputBuilder) WithTextStyle(style lipgloss.Color) *InputBuilder {
	ib.model.TextStyle = constants.GetAnsiStyle(style)
	return ib
}

// Build returns the configured textinput.Model
func (ib *InputBuilder) Build() textinput.Model {
	return ib.model
}

// Common input constructors for convenience

// NewBranchInput creates a styled input for branch names
func NewBranchInput() textinput.Model {
	return NewInput().
		WithPlaceholder("Branch").
		WithCharLimit(256).
		WithWidth(20).
		WithCursorStyle(constants.Yellow).
		Build()
}

// NewUserInput creates a styled input for usernames
func NewUserInput() textinput.Model {
	return NewInput().
		WithPlaceholder("User").
		WithCharLimit(32).
		WithWidth(50).
		WithCursorStyle(constants.Yellow).
		Build()
}

// NewTokenInput creates a styled input for tokens/passwords
func NewTokenInput() textinput.Model {
	return NewInput().
		WithPlaceholder("Token").
		WithCharLimit(256).
		WithWidth(120).
		WithCursorStyle(constants.Yellow).
		Build()
}

// NewEmailInput creates a styled input for email addresses
func NewEmailInput() textinput.Model {
	return NewInput().
		WithPlaceholder("Email").
		WithCharLimit(100).
		WithWidth(60).
		WithCursorStyle(constants.Yellow).
		Build()
}

// NewNameInput creates a styled input for names
func NewNameInput() textinput.Model {
	return NewInput().
		WithPlaceholder("Full Name").
		WithCharLimit(100).
		WithWidth(60).
		WithCursorStyle(constants.Yellow).
		Build()
}

// NewSigningKeyInput creates a styled input for GPG signing keys
func NewSigningKeyInput() textinput.Model {
	return NewInput().
		WithPlaceholder("Signing Key (GPG)").
		WithCharLimit(100).
		WithWidth(60).
		WithCursorStyle(constants.Yellow).
		Build()
}

// NewGenericInput creates a generic text input with custom placeholder
func NewGenericInput(placeholder string) textinput.Model {
	return NewInput().
		WithPlaceholder(placeholder).
		Build()
}
