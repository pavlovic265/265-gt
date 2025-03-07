package components

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type YesNoPrompt struct {
	question  string
	choices   []string
	current   int
	Decisions map[string]bool
	Quitting  bool
}

func NewYesNoPrompt(question string, choices []string) YesNoPrompt {
	return YesNoPrompt{
		question:  question,
		choices:   choices,
		current:   0,
		Decisions: make(map[string]bool),
	}
}

func (m YesNoPrompt) Init() tea.Cmd {
	return nil
}

func (m YesNoPrompt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "y", "Y", "enter":
			m.Decisions[m.choices[m.current]] = true
			m.current++
		case "n", "N":
			m.Decisions[m.choices[m.current]] = false
			m.current++
		case "q", "ctrl+c", "esc":
			m.Quitting = true
			return m, tea.Quit
		}

		if m.current >= len(m.choices) {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m YesNoPrompt) View() string {
	if m.Quitting {
		return "Canceled by user\n"
	}
	if m.current >= len(m.choices) {
		var output strings.Builder
		for branch, decision := range m.Decisions {
			output.WriteString(fmt.Sprintf("Deleted %s: %t\n", branch, decision))
		}
		return output.String()
	}
	currentBranch := m.choices[m.current]
	prompt := fmt.Sprintf(m.question, currentBranch)

	return prompt
}
