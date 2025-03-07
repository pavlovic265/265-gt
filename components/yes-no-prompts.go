package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type YesNoPrompt struct {
	question string
	answer   string
	Quitting bool
}

func NewYesNoPrompt(question string) YesNoPrompt {
	return YesNoPrompt{
		question: question,
	}
}

func (m YesNoPrompt) Init() tea.Cmd {
	return nil
}

func (m YesNoPrompt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "y", "Y":
			m.answer = "Yes"
			return m, tea.Quit
		case "n", "N":
			m.answer = "No"
			return m, tea.Quit
		case "enter":
			m.answer = "Yes"
			return m, tea.Quit
		case "q", "ctrl+c", "esc":
			m.Quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m YesNoPrompt) View() string {
	if m.Quitting {
		return "Canceled by user\n"
	}
	return fmt.Sprintf("%s %s", m.question, m.answer)
}
