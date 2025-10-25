package createconfig

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
)

type accountsModel struct {
	focusIndex    int
	inputs        []textinput.Model
	accounts      []config.Account
	platform      constants.Platform
	platformIndex int
}

func newAccountsModel() accountsModel {
	accountsModel := accountsModel{
		inputs: make([]textinput.Model, 2),
	}

	accountsModel.inputs[0] = components.NewUserInput()
	accountsModel.inputs[1] = components.NewTokenInput()
	accountsModel.focusIndex = 0
	accountsModel.platform = constants.GitHubPlatform // Default to GitHub
	accountsModel.platformIndex = 0                   // Default to first platform (GitHub)

	return accountsModel
}

var platformOptions = []constants.Platform{
	constants.GitHubPlatform,
	constants.GitLabPlatform,
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
				if am.focusIndex == len(am.inputs)+1 {
					return am.handleDone()
					// If focus is on Add button.
				} else if am.focusIndex == len(am.inputs)+2 {
					return am.handleAdd()
				}
			}

			// Handle platform selection with Tab
			if am.focusIndex == len(am.inputs) { // Platform selection area
				if key == tea.KeyTab.String() {
					return am.handlePlatformSelection(key)
				}
			}

			// Cycle indexes (but skip Tab when in platform selection)
			if am.focusIndex == len(am.inputs) && key == tea.KeyTab.String() {
				// Don't cycle, let platform selection handle it
				return am, nil
			}
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

	// Add platform selection
	b.WriteString("\nPlatform: ")
	for i, platform := range platformOptions {
		platformStr := platform.String()
		if i == am.platformIndex {
			if am.focusIndex == len(am.inputs) {
				b.WriteString("(• " + platformStr + ")")
			} else {
				b.WriteString("(•) " + platformStr)
			}
		} else {
			if am.focusIndex == len(am.inputs) {
				b.WriteString("( ) " + platformStr)
			} else {
				b.WriteString("( ) " + platformStr)
			}
		}
		if i < len(platformOptions)-1 {
			b.WriteString("  ")
		}
	}

	doneButton := components.NewDoneButton(am.focusIndex == len(am.inputs)+1)
	addButton := components.NewAddButton(am.focusIndex == len(am.inputs)+2)

	fmt.Fprintf(&b, "\n\n%s  %s\n\n", doneButton.Render(), addButton.Render())

	return b.String()
}

func (am accountsModel) handleDone() (tea.Model, tea.Cmd) {
	if am.inputs[0].Value() != "" && am.inputs[1].Value() != "" {
		platform := platformOptions[am.platformIndex]
		am.accounts = append(am.accounts, config.Account{
			User:     am.inputs[0].Value(),
			Token:    am.inputs[1].Value(),
			Platform: platform,
		})
	}
	return am, tea.Quit
}

func (am accountsModel) handleAdd() (tea.Model, tea.Cmd) {
	platform := platformOptions[am.platformIndex]
	am.accounts = append(am.accounts, config.Account{
		User:     am.inputs[0].Value(),
		Token:    am.inputs[1].Value(),
		Platform: platform,
	})

	am.inputs[0] = components.NewUserInput()
	am.inputs[1] = components.NewTokenInput()
	am.focusIndex = 0

	cmds := make([]tea.Cmd, len(am.inputs))
	return am, tea.Batch(cmds...)
}

func (am accountsModel) handlePlatformSelection(key string) (tea.Model, tea.Cmd) {
	// Cycle through platform options with Tab
	am.platformIndex++
	if am.platformIndex >= len(platformOptions) {
		am.platformIndex = 0
	}
	am.platform = platformOptions[am.platformIndex]
	return am, nil
}

func (am accountsModel) handleCycle(key string) (tea.Model, tea.Cmd) {
	if key == tea.KeyUp.String() || key == tea.KeyShiftTab.String() || key == tea.KeyCtrlK.String() {
		am.focusIndex--
	} else {
		am.focusIndex++
	}

	// Wrap focusIndex to range [0, len(am.inputs)+2] (inputs + platform + buttons).
	if am.focusIndex > len(am.inputs)+2 {
		am.focusIndex = 0
	} else if am.focusIndex < 0 {
		am.focusIndex = len(am.inputs) + 2
	}

	cmds := make([]tea.Cmd, len(am.inputs))
	for i := 0; i <= len(am.inputs)-1; i++ {
		if i == am.focusIndex {
			// Set focused state
			cmds[i] = am.inputs[i].Focus()
			am.inputs[i].PromptStyle = lipgloss.NewStyle()
			am.inputs[i].TextStyle = lipgloss.NewStyle()
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
