package createconfig

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/config"
)

var (
	DoneButtonFocus = "[ Done ]"
	DoneButtonBlur  = "[ Done ]"

	AddButtonFocus = "[ Add ]"
	AddButtonBlur  = "[ Add ]"
)

type accountsModel struct {
	focusIndex int
	inputs     []textinput.Model
	accounts   []config.Account
}

func newAccountsModel() accountsModel {
	accountsModel := accountsModel{
		inputs: make([]textinput.Model, 2),
	}

	accountsModel.inputs[0] = buildUserInput()
	accountsModel.inputs[1] = buildTokenInput()
	accountsModel.focusIndex = 0

	return accountsModel
}

func buildUserInput() textinput.Model {
	t := textinput.New()

	t.Cursor.Style = config.GetSuccessStyle()

	t.Placeholder = "User"
	t.Focus()
	t.CharLimit = 32
	t.PromptStyle = config.GetSuccessStyle()

	t.TextStyle = config.GetSuccessStyle()
	return t
}

func buildTokenInput() textinput.Model {
	t := textinput.New()

	t.Cursor.Style = config.GetSuccessStyle()

	t.Placeholder = "Token"
	t.CharLimit = 128
	t.PromptStyle = config.GetSuccessStyle()
	t.TextStyle = config.GetSuccessStyle()
	return t
}

func (am accountsModel) Init() tea.Cmd {
	return textinput.Blink
}

func (am accountsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String(), tea.KeyEsc.String(), tea.KeyCtrlQ.String():
			return am, tea.Quit

		// Set focus to next input
		case tea.KeyTab.String(), tea.KeyShiftTab.String(),
			tea.KeyEnter.String(), tea.KeyUp.String(), tea.KeyDown.String(),
			tea.KeyCtrlJ.String(), tea.KeyCtrlK.String():
			key := msg.String()

			// Handle Enter key separately for buttons.
			if key == tea.KeyEnter.String() {
				// If focus is on Done button.
				if am.focusIndex == len(am.inputs) {
					return am.handleDone()
					// If focus is on Add button.
				} else if am.focusIndex == len(am.inputs)+1 {
					return am.handleAdd()
				}
			}

			// Cycle indexes
			return am.handleCycle(key)
		}
	}

	// Handle character input and blinking
	cmds := make([]tea.Cmd, len(am.inputs))
	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range am.inputs {
		am.inputs[i], cmds[i] = am.inputs[i].Update(msg)
	}

	return am, tea.Batch(cmds...)
}

func (am accountsModel) View() string {
	var b strings.Builder

	for i := range am.inputs {
		b.WriteString(am.inputs[i].View())
		if i < len(am.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	doneButton := DoneButtonBlur
	if am.focusIndex == len(am.inputs) {
		doneButton = config.GetSuccessStyle().Render(DoneButtonFocus)
	} else {
		doneButton = config.GetDebugStyle().Render(DoneButtonBlur)
	}

	addButton := AddButtonBlur
	if am.focusIndex == len(am.inputs)+1 {
		addButton = config.GetSuccessStyle().Render(AddButtonFocus)
	} else {
		addButton = config.GetDebugStyle().Render(AddButtonBlur)
	}

	fmt.Fprintf(&b, "\n\n%s  %s\n\n", doneButton, addButton)

	return b.String()
}

func (am accountsModel) handleDone() (tea.Model, tea.Cmd) {
	if am.inputs[0].Value() != "" && am.inputs[1].Value() != "" {
		am.accounts = append(am.accounts, config.Account{User: am.inputs[0].Value(), Token: am.inputs[1].Value()})
	}
	return am, tea.Quit
}

func (am accountsModel) handleAdd() (tea.Model, tea.Cmd) {
	am.accounts = append(am.accounts, config.Account{User: am.inputs[0].Value(), Token: am.inputs[1].Value()})

	am.inputs[0] = buildUserInput()
	am.inputs[1] = buildTokenInput()
	am.focusIndex = 0

	cmds := make([]tea.Cmd, len(am.inputs))
	return am, tea.Batch(cmds...)
}

func (am accountsModel) handleCycle(key string) (tea.Model, tea.Cmd) {
	if key == tea.KeyUp.String() || key == tea.KeyShiftTab.String() || key == tea.KeyCtrlK.String() {
		am.focusIndex--
	} else {
		am.focusIndex++
	}

	// Wrap focusIndex to range [0, len(am.inputs)+1].
	if am.focusIndex > len(am.inputs)+1 {
		am.focusIndex = 0
	} else if am.focusIndex < 0 {
		am.focusIndex = len(am.inputs) + 1
	}

	cmds := make([]tea.Cmd, len(am.inputs))
	for i := 0; i <= len(am.inputs)-1; i++ {
		if i == am.focusIndex {
			// Set focused state
			cmds[i] = am.inputs[i].Focus()
			am.inputs[i].PromptStyle = config.GetSuccessStyle()
			am.inputs[i].TextStyle = config.GetSuccessStyle()
			continue
		}
		// Remove focused state
		am.inputs[i].Blur()
		am.inputs[i].PromptStyle = lipgloss.NewStyle()
		am.inputs[i].TextStyle = lipgloss.NewStyle()
	}

	return am, tea.Batch(cmds...)
}

func HandleAddAccunts() ([]config.Account, error) {
	selectModel := newAccountsModel()
	selectProgram := tea.NewProgram(selectModel)
	m, err := selectProgram.Run()
	if err != nil {
		return nil, err
	}

	if m, ok := m.(accountsModel); ok {
		return m.accounts, nil
	}

	return nil, fmt.Errorf("faild to read accounts")
}
