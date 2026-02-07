package account

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/components"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/utils/log"
)

type tokenSetupModel struct {
	account    *config.Account
	tokenInput textinput.Model
	focusIndex int // 0 = input, 1 = skip button, 2 = save button
	skipped    bool
	saved      bool
}

func newTokenSetupModel(account *config.Account) tokenSetupModel {
	tokenInput := components.NewTokenInput()
	tokenInput.Focus()

	// Prefill if account already has a token
	if account.Token != "" {
		tokenInput.SetValue(account.Token)
	}

	return tokenSetupModel{
		account:    account,
		tokenInput: tokenInput,
		focusIndex: 0,
	}
}

func (m tokenSetupModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m tokenSetupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String(), tea.KeyCtrlQ.String(), tea.KeyEsc.String():
			m.skipped = true
			return m, tea.Quit

		case tea.KeyTab.String(), tea.KeyShiftTab.String():
			if msg.String() == tea.KeyShiftTab.String() {
				m.focusIndex--
				if m.focusIndex < 0 {
					m.focusIndex = 2
				}
			} else {
				m.focusIndex++
				if m.focusIndex > 2 {
					m.focusIndex = 0
				}
			}
			m.updateFocus()
			return m, nil

		case tea.KeyEnter.String():
			if m.focusIndex == 1 {
				// Skip button
				m.skipped = true
				return m, tea.Quit
			}
			if m.focusIndex == 2 {
				// Save button
				m.saved = true
				return m, tea.Quit
			}
			// If on input, move to next
			m.focusIndex++
			m.updateFocus()
			return m, nil
		}
	}

	if m.focusIndex == 0 {
		var cmd tea.Cmd
		m.tokenInput, cmd = m.tokenInput.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m *tokenSetupModel) updateFocus() {
	if m.focusIndex == 0 {
		m.tokenInput.Focus()
		m.tokenInput.PromptStyle = constants.GetSuccessAnsiStyle()
		m.tokenInput.TextStyle = constants.GetAnsiStyle(constants.White)
	} else {
		m.tokenInput.Blur()
		m.tokenInput.PromptStyle = constants.GetAnsiStyle(constants.BrightBlack)
		m.tokenInput.TextStyle = constants.GetAnsiStyle(constants.BrightBlack)
	}
}

func (m tokenSetupModel) View() string {
	var b strings.Builder

	titleStyle := constants.GetWarningAnsiStyle().Bold(true)
	dimStyle := constants.GetAnsiStyle(constants.BrightBlack)
	highlightStyle := constants.GetWarningAnsiStyle()
	infoStyle := constants.GetAnsiStyle(constants.Cyan)

	b.WriteString(titleStyle.Render(fmt.Sprintf("Token Setup for %s (%s)", m.account.User, m.account.Platform)))
	b.WriteString("\n\n")

	// Platform-specific instructions
	b.WriteString(infoStyle.Render("To create a token, go to:"))
	b.WriteString("\n")
	if m.account.Platform == constants.GitHubPlatform {
		b.WriteString("  GitHub -> Settings -> Developer settings -> Personal access tokens\n")
	} else {
		b.WriteString("  GitLab -> Preferences -> Access Tokens\n")
	}
	b.WriteString("\n")

	// Limitations warning
	b.WriteString(dimStyle.Render("Without a token you cannot:"))
	b.WriteString("\n")
	b.WriteString(dimStyle.Render("  - Create pull/merge requests"))
	b.WriteString("\n")
	b.WriteString(dimStyle.Render("  - View PR/MR status"))
	b.WriteString("\n")
	b.WriteString(dimStyle.Render("  - Access private repos via API"))
	b.WriteString("\n\n")

	// Token input
	if m.focusIndex == 0 {
		b.WriteString(highlightStyle.Render(">") + " ")
	} else {
		b.WriteString("  ")
	}
	b.WriteString(m.tokenInput.View())
	b.WriteString("\n\n")

	// Buttons
	skipButton := components.NewSkipButton(m.focusIndex == 1)
	saveButton := components.NewSaveButton(m.focusIndex == 2)
	b.WriteString(fmt.Sprintf("%s  %s\n", skipButton.Render(), saveButton.Render()))

	b.WriteString("\n")
	b.WriteString(dimStyle.Render("You can add a token later with: "))
	b.WriteString(highlightStyle.Render("gt account edit -t"))

	b.WriteString("\n\n")
	b.WriteString(dimStyle.Render("Press "))
	b.WriteString(highlightStyle.Render("Ctrl+Q"))
	b.WriteString(dimStyle.Render(" to quit"))

	return b.String()
}

func HandleTokenSetup(account *config.Account) error {
	model := newTokenSetupModel(account)
	program := tea.NewProgram(model)
	m, err := program.Run()
	if err != nil {
		return log.Error("failed to run token setup", err)
	}

	result, ok := m.(tokenSetupModel)
	if !ok {
		return nil
	}

	if result.skipped {
		log.Warning("Token setup skipped. Some features will be limited.")
		return nil
	}

	if result.saved {
		token := result.tokenInput.Value()
		if token != "" {
			account.Token = token
			log.Success("Token saved successfully!")
		} else {
			log.Warning("No token provided. Some features will be limited.")
		}
	}

	return nil
}
