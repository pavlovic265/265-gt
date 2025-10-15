package components

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/constants"
)

type YesNoPrompt struct {
	question string
	answer   string
	Quitting bool
}

// Styling definitions for yes/no prompt
var (
	// Question style
	questionStyle = lipgloss.NewStyle().
			Foreground(constants.Blue)

	// Options style
	optionsStyle = lipgloss.NewStyle().
			Foreground(constants.BrightBlack)

	// Key styles
	yesKeyStyle = lipgloss.NewStyle().
			Foreground(constants.Green).
			Bold(true)

	noKeyStyle = lipgloss.NewStyle().
			Foreground(constants.Red).
			Bold(true)

	enterKeyStyle = lipgloss.NewStyle().
			Foreground(constants.Yellow).
			Bold(true)

	quitKeyStyle = lipgloss.NewStyle().
			Foreground(constants.Yellow).
			Bold(true)

	// Canceled message style
	canceledStyle = lipgloss.NewStyle().
			Foreground(constants.Red)
)

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
			constants.ErrorIcon,
			canceledStyle.Render("Operation canceled by user"))
	}

	var content strings.Builder

	// Question
	content.WriteString(questionStyle.Render(m.question))
	content.WriteString("\n")

	// Options with styled keys
	content.WriteString(optionsStyle.Render("Press "))
	content.WriteString(yesKeyStyle.Render("Yes (Y)"))
	content.WriteString(optionsStyle.Render(", "))
	content.WriteString(noKeyStyle.Render("No (N)"))
	content.WriteString(optionsStyle.Render(", "))
	content.WriteString(enterKeyStyle.Render("ENTER (Yes)"))
	content.WriteString(optionsStyle.Render(", or "))
	content.WriteString(quitKeyStyle.Render("Ctrl+Q"))
	content.WriteString(optionsStyle.Render(" to quit"))

	return content.String()
}
