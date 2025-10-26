package account

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

// Styling definitions for account form interface
var (
	optionsStyle = lipgloss.NewStyle().
			Foreground(constants.BrightBlack)
	quitKeyStyle = lipgloss.NewStyle().
			Foreground(constants.Yellow).
			Bold(true)
	cursorStyle = lipgloss.NewStyle().
			Foreground(constants.Yellow)
	selectedDotStyle = lipgloss.NewStyle().
				Foreground(constants.Yellow)
)

type accountsModel struct {
	focusIndex    int
	inputs        []textinput.Model
	accounts      []config.Account
	platform      constants.Platform
	platformIndex int
	editMode      bool // true = edit mode (only Done button), false = add mode (Done + Add buttons)
}

func newAccountsModel() accountsModel {
	return newAccountsModelWithData(nil)
}

func newAccountsModelWithData(account *config.Account) accountsModel {
	accountsModel := accountsModel{
		inputs:   make([]textinput.Model, 4),
		editMode: account != nil, // Set edit mode if account is provided
	}

	accountsModel.inputs[0] = components.NewUserInput()
	accountsModel.inputs[1] = components.NewTokenInput()
	accountsModel.inputs[2] = components.NewEmailInput()
	accountsModel.inputs[3] = components.NewNameInput()
	accountsModel.focusIndex = 0
	accountsModel.platform = constants.GitHubPlatform // Default to GitHub
	accountsModel.platformIndex = 0                   // Default to first platform (GitHub)

	// Prefill if account data is provided
	if account != nil {
		accountsModel.inputs[0].SetValue(account.User)
		accountsModel.inputs[1].SetValue(account.Token)
		accountsModel.inputs[2].SetValue(account.Email)
		accountsModel.inputs[3].SetValue(account.Name)

		// Set platform
		accountsModel.platform = account.Platform
		for i, platform := range platformOptions {
			if platform == account.Platform {
				accountsModel.platformIndex = i
				break
			}
		}
	}

	return accountsModel
}

var platformOptions = []constants.Platform{
	constants.GitHubPlatform,
	constants.GitLabPlatform,
}

func (am accountsModel) Init() tea.Cmd {
	am.updateFocus()
	return textinput.Blink
}

func (am accountsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String(), tea.KeyEsc.String(), tea.KeyCtrlQ.String():
			return am, tea.Quit

		case tea.KeyTab.String(), tea.KeyShiftTab.String():
			key := msg.String()

			// Tab always cycles through fields
			return am.handleCycle(key)

		case tea.KeyUp.String(), tea.KeyDown.String(),
			tea.KeyCtrlJ.String(), tea.KeyCtrlK.String():
			key := msg.String()

			// When on platform, arrow keys toggle between GitHub/GitLab
			if am.focusIndex == len(am.inputs) {
				return am.handlePlatformSelection(key)
			}

			// Otherwise cycle through fields
			return am.handleCycle(key)

		case "j", "k":
			// Only handle j/k when on platform field (not in text inputs)
			if am.focusIndex == len(am.inputs) {
				return am.handlePlatformSelection(msg.String())
			}
			// Otherwise let it pass through to text input

		case tea.KeyEnter.String():
			// Handle Enter key for buttons
			if am.focusIndex == len(am.inputs)+1 {
				// Done button
				return am.handleDone()
			} else if !am.editMode && am.focusIndex == len(am.inputs)+2 {
				// Add button (only in add mode)
				return am.handleAdd()
			} else if am.focusIndex == len(am.inputs) {
				// If focus is on platform field, move to next element
				return am.handleCycle(tea.KeyTab.String())
			} else if am.focusIndex < len(am.inputs) {
				// If focus is on input field, cycle to next element
				return am.handleCycle(tea.KeyTab.String())
			}
		}
	}

	// Only update the text input if focus is on it
	if am.focusIndex < len(am.inputs) {
		am.inputs[am.focusIndex], cmd = am.inputs[am.focusIndex].Update(msg)
	}

	return am, cmd
}

func (am accountsModel) View() string {
	var b strings.Builder

	for i := range am.inputs {
		// Add cursor indicator for focused input
		if am.focusIndex == i {
			b.WriteString(cursorStyle.Render(">") + " ")
		} else {
			b.WriteString("  ")
		}
		b.WriteString(am.inputs[i].View())
		if i < len(am.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	// Add platform selection
	b.WriteString("\n")
	if am.focusIndex == len(am.inputs) {
		b.WriteString(cursorStyle.Render(">") + " ")
	} else {
		b.WriteString("  ")
	}
	b.WriteString("Platform: ")
	for i, platform := range platformOptions {
		platformStr := platform.String()
		if i == am.platformIndex {
			b.WriteString(selectedDotStyle.Render("(•)") + " " + platformStr)
		} else {
			b.WriteString("( ) " + platformStr)
		}
		if i < len(platformOptions)-1 {
			b.WriteString("  ")
		}
	}
	b.WriteString("\n")

	doneButton := components.NewDoneButton(am.focusIndex == len(am.inputs)+1)

	if am.editMode {
		// Edit mode: only show Done button
		fmt.Fprintf(&b, "\n%s\n\n", doneButton.Render())
	} else {
		// Add mode: show Done and Add buttons
		addButton := components.NewAddButton(am.focusIndex == len(am.inputs)+2)
		fmt.Fprintf(&b, "\n%s  %s\n\n", doneButton.Render(), addButton.Render())
	}

	// Add quit instruction
	b.WriteString(optionsStyle.Render("Press "))
	b.WriteString(quitKeyStyle.Render("Ctrl+Q"))
	b.WriteString(optionsStyle.Render(" to quit"))

	return b.String()
}

func (am accountsModel) handleDone() (tea.Model, tea.Cmd) {
	if am.inputs[0].Value() != "" && am.inputs[1].Value() != "" {
		platform := platformOptions[am.platformIndex]
		am.accounts = append(am.accounts, config.Account{
			User:     am.inputs[0].Value(),
			Token:    am.inputs[1].Value(),
			Platform: platform,
			Email:    am.inputs[2].Value(),
			Name:     am.inputs[3].Value(),
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
		Email:    am.inputs[2].Value(),
		Name:     am.inputs[3].Value(),
	})

	am.inputs[0] = components.NewUserInput()
	am.inputs[1] = components.NewTokenInput()
	am.inputs[2] = components.NewEmailInput()
	am.inputs[3] = components.NewNameInput()
	am.focusIndex = 0

	cmds := make([]tea.Cmd, len(am.inputs))
	return am, tea.Batch(cmds...)
}

func (am accountsModel) handlePlatformSelection(key string) (tea.Model, tea.Cmd) {
	// Cycle through platform options
	if key == tea.KeyUp.String() || key == tea.KeyCtrlK.String() || key == "k" {
		// Move backward (Up, k)
		am.platformIndex--
		if am.platformIndex < 0 {
			am.platformIndex = len(platformOptions) - 1
		}
	} else {
		// Move forward (Down, j)
		am.platformIndex++
		if am.platformIndex >= len(platformOptions) {
			am.platformIndex = 0
		}
	}
	am.platform = platformOptions[am.platformIndex]
	return am, nil
}

func (am accountsModel) updateFocus() {
	// Handle focus based on focusIndex - only focus the active input
	for i := range am.inputs {
		if i == am.focusIndex {
			am.inputs[i].Focus()
			am.inputs[i].PromptStyle = constants.GetSuccessAnsiStyle()
			am.inputs[i].TextStyle = constants.GetAnsiStyle(constants.White)
		} else {
			am.inputs[i].Blur()
			am.inputs[i].PromptStyle = lipgloss.NewStyle()
			am.inputs[i].TextStyle = lipgloss.NewStyle()
		}
	}
}

func (am accountsModel) handleCycle(key string) (tea.Model, tea.Cmd) {
	if key == tea.KeyUp.String() || key == tea.KeyShiftTab.String() || key == tea.KeyCtrlK.String() {
		am.focusIndex--
	} else {
		am.focusIndex++
	}

	// Calculate max focus index based on mode
	// Add mode: inputs + platform + Done button + Add button = len(am.inputs) + 2
	// Edit mode: inputs + platform + Done button = len(am.inputs) + 1
	maxIndex := len(am.inputs) + 1 // Done button (edit mode)
	if !am.editMode {
		maxIndex = len(am.inputs) + 2 // Done + Add buttons (add mode)
	}

	// Wrap focusIndex
	if am.focusIndex > maxIndex {
		am.focusIndex = 0
	} else if am.focusIndex < 0 {
		am.focusIndex = maxIndex
	}

	// Update focus state after changing focusIndex
	am.updateFocus()

	return am, nil
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

func HandleEditAccount(account *config.Account) (*config.Account, error) {
	selectModel := newAccountsModelWithData(account)
	selectProgram := tea.NewProgram(selectModel)
	m, err := selectProgram.Run()
	if err != nil {
		return nil, err
	}

	if m, ok := m.(accountsModel); ok {
		if len(m.accounts) > 0 {
			return &m.accounts[0], nil
		}
	}

	return nil, fmt.Errorf("failed to edit account")
}
