package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/constants"
)

type YesNoPrompt struct {
	question string
	answer   string
	Quitting bool
}

func (m YesNoPrompt) IsYes() bool {
	if m.answer != "" && m.answer == "Yes" {
		return true
	}
	return false
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
		case tea.KeyEnter.String():
			m.answer = "Yes"
			return m, tea.Quit
		case tea.KeyEsc.String(), tea.KeyCtrlC.String(), tea.KeyCtrlQ.String():
			m.Quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m YesNoPrompt) View() string {
	if m.Quitting {
		return fmt.Sprintf("\n%s %s\n",
			constants.ErrorIconOnly(),
			constants.GetErrorStyle().Render("Operation canceled by user"))
	}

	// Style the question with options
	s := fmt.Sprintf("%s\n", constants.GetInfoStyle().Render(m.question))
	s += fmt.Sprintf("%s %s %s %s\n",
		constants.GetDebugStyle().Render("Press"),
		constants.GetSuccessStyle().Render("Y"),
		constants.GetDebugStyle().Render("for Yes,"),
		constants.GetErrorStyle().Render("N"))
	s += fmt.Sprintf("%s %s %s\n",
		constants.GetDebugStyle().Render("for No, or"),
		constants.GetCommandStyle().Render("ENTER"),
		constants.GetDebugStyle().Render("for Yes."))

	return s
}
