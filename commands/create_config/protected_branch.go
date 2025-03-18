package createconfig

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func HandleAddProtectedBranch() ([]string, error) {
	protectedBranch := newProtectedBranchModel()
	protectedBranchProgram := tea.NewProgram(protectedBranch)
	m, err := protectedBranchProgram.Run()
	if err != nil {
		return nil, err
	}

	if m, ok := m.(protectedBranchModele); ok {
		return m.branches, nil
	}

	return nil, fmt.Errorf("faild to read accounts")
}

type protectedBranchModele struct {
	focusIndex int
	branch     textinput.Model
	branches   []string
}

func newProtectedBranchModel() protectedBranchModele {
	return protectedBranchModele{
		branch:     buildTextInput(),
		focusIndex: 0,
	}
}

func buildTextInput() textinput.Model {
	pbm := textinput.New()
	pbm.Placeholder = "Branch"
	pbm.Focus()
	pbm.CharLimit = 256
	pbm.Width = 20

	return pbm
}

func (m protectedBranchModele) Init() tea.Cmd {
	return textinput.Blink
}

func (pbm protectedBranchModele) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String(), tea.KeyEsc.String(), tea.KeyCtrlQ.String():
			return pbm, tea.Quit
		case tea.KeyTab.String(), tea.KeyShiftTab.String(),
			tea.KeyEnter.String(), tea.KeyUp.String(), tea.KeyDown.String(),
			tea.KeyCtrlJ.String(), tea.KeyCtrlK.String():
			key := msg.String()

			// Handle Enter key separately for buttons.
			if key == tea.KeyEnter.String() {
				switch pbm.focusIndex {
				case 1:
					return pbm.handleDone()
				case 2:
					return pbm.handleAdd()
				}
			}

			// Cycle indexes
			return pbm.handleCycle(key)
		}
	}

	pbm.branch, cmd = pbm.branch.Update(msg)
	return pbm, cmd
}

func (pbm protectedBranchModele) handleDone() (tea.Model, tea.Cmd) {
	if pbm.branch.Value() != "" {
		pbm.branches = append(pbm.branches, pbm.branch.Value())
	}
	return pbm, tea.Quit
}

func (pbm protectedBranchModele) handleAdd() (tea.Model, tea.Cmd) {
	pbm.branches = append(pbm.branches, pbm.branch.Value())

	pbm.branch = buildTextInput()
	pbm.focusIndex = 0

	return pbm, nil
}

func (pbm protectedBranchModele) handleCycle(key string) (tea.Model, tea.Cmd) {
	if key == tea.KeyUp.String() || key == tea.KeyShiftTab.String() || key == tea.KeyCtrlK.String() {
		pbm.focusIndex--
	} else {
		pbm.focusIndex++
	}

	if pbm.focusIndex > 2 {
		pbm.focusIndex = 0
	} else if pbm.focusIndex < 0 {
		pbm.focusIndex = 0
	}

	return pbm, nil
}

func (pbm protectedBranchModele) View() string {
	var b strings.Builder
	b.WriteString(pbm.branch.View())
	b.WriteRune('\n')

	doneButton := &DoneButtonBlur
	if pbm.focusIndex == 1 {
		doneButton = &DoneButtonFocus
	}
	addButton := &AddButtonBlur
	if pbm.focusIndex == 2 {
		addButton = &AddButtonFocus
	}
	fmt.Fprintf(&b, "\n%s  %s\n\n", *doneButton, *addButton)

	return b.String()
}
