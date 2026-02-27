package components

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// NewStringListModel creates a ListModel for string choices with default formatter and matcher.
func NewStringListModel(choices []string) ListModel[string] {
	return ListModel[string]{
		AllChoices: choices,
		Choices:    choices,
		Cursor:     0,
		Query:      "",
		Formatter:  func(s string) string { return s },
		Matcher:    func(s, query string) bool { return strings.Contains(strings.ToLower(s), strings.ToLower(query)) },
	}
}

// SelectString displays a selection list and returns the selected string.
// Returns empty string if user cancelled or no selection was made.
func SelectString(choices []string) (string, error) {
	if len(choices) == 0 {
		return "", nil
	}

	model := NewStringListModel(choices)
	program := tea.NewProgram(model)

	finalModel, err := program.Run()
	if err != nil {
		return "", err
	}

	if m, ok := finalModel.(ListModel[string]); ok {
		return m.Selected, nil
	}

	return "", nil
}
