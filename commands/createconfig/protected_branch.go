package createconfig

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/utils/log"
)

var (
	optionsStyle = lipgloss.NewStyle().
			Foreground(constants.BrightBlack)
	quitKeyStyle = lipgloss.NewStyle().
			Foreground(constants.Yellow).
			Bold(true)
)

func HandleAddProtectedBranch() ([]string, error) {
	protectedBranch := newProtectedBranchModel()
	protectedBranchProgram := tea.NewProgram(&protectedBranch)
	m, err := protectedBranchProgram.Run()
	if err != nil {
		return nil, log.Error("Failed to run protected branch interface", err)
	}

	if m, ok := m.(protectedBranchModele); ok {
		// If user quit without clicking quitting, return error
		if m.quitting {
			return nil, nil
		}
		return m.branches, nil
	}

	return nil, log.ErrorMsg("Failed to read protected branch configuration")
}

type protectedBranchModele struct {
	focusIndex int
	branch     textinput.Model
	branches   []string
	quitting   bool
}

func newProtectedBranchModel() protectedBranchModele {
	return protectedBranchModele{
		branch:     components.NewBranchInput(),
		focusIndex: 0,
		quitting:   false,
	}
}

func (m protectedBranchModele) Init() tea.Cmd {
	return textinput.Blink
}

func (m protectedBranchModele) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// Handle focus based on focusIndex
	if m.focusIndex == 0 {
		m.branch.Focus()
	} else {
		m.branch.Blur()
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String(), tea.KeyEsc.String(), tea.KeyCtrlQ.String():
			m.quitting = true
			return m, tea.Quit
		case tea.KeyTab.String(), tea.KeyShiftTab.String(),
			tea.KeyUp.String(), tea.KeyDown.String(),
			tea.KeyCtrlJ.String(), tea.KeyCtrlK.String():
			key := msg.String()
			// Cycle indexes
			return m.handleCycle(key)
		case tea.KeyEnter.String():
			// Handle Enter key for buttons only
			switch m.focusIndex {
			case 1:
				return m.handlequitting()
			case 2:
				return m.handleAdd()
			default:
				// If focus is on input field, cycle to next element
				return m.handleCycle(tea.KeyTab.String())
			}
		}
	}

	// Only update the text input if focus is on it (index 0)
	if m.focusIndex == 0 {
		m.branch, cmd = m.branch.Update(msg)
	}
	return m, cmd
}

func (m protectedBranchModele) handlequitting() (tea.Model, tea.Cmd) {
	if m.branch.Value() != "" {
		m.branches = append(m.branches, m.branch.Value())
	}
	return m, tea.Quit
}

func (m protectedBranchModele) handleAdd() (tea.Model, tea.Cmd) {
	m.branches = append(m.branches, m.branch.Value())

	m.branch = components.NewBranchInput()
	m.focusIndex = 0

	return m, nil
}

func (m protectedBranchModele) handleCycle(key string) (tea.Model, tea.Cmd) {
	if key == tea.KeyUp.String() || key == tea.KeyShiftTab.String() || key == tea.KeyCtrlK.String() {
		m.focusIndex--
	} else {
		m.focusIndex++
	}

	// Proper cycling: 0 (input) -> 1 (quitting) -> 2 (add) -> 0 (input)
	if m.focusIndex > 2 {
		m.focusIndex = 0
	} else if m.focusIndex < 0 {
		m.focusIndex = 2
	}

	return m, nil
}

func (m protectedBranchModele) View() string {
	var b strings.Builder

	b.WriteString(m.branch.View())
	b.WriteRune('\n')

	doneButton := components.NewDoneButton(m.focusIndex == 1)
	addButton := components.NewAddButton(m.focusIndex == 2)

	fmt.Fprintf(&b, "\n%s  %s\n\n", doneButton.Render(), addButton.Render())

	// Add quit instruction
	b.WriteString(optionsStyle.Render("Press "))
	b.WriteString(quitKeyStyle.Render("Ctrl+Q"))
	b.WriteString(optionsStyle.Render(" to quit"))

	return b.String()
}
