package account

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/ui/components"
	"github.com/pavlovic265/265-gt/ui/theme"
	"github.com/pavlovic265/265-gt/utils/log"
)

type gpgSetupModel struct {
	account    *config.Account
	gpgInput   textinput.Model
	focusIndex int // 0 = input, 1 = skip button, 2 = save button
	skipped    bool
	saved      bool
}

func newGPGSetupModel(account *config.Account) gpgSetupModel {
	gpgInput := components.NewSigningKeyInput()
	gpgInput.Focus()

	// Prefill if account already has a signing key
	if account.SigningKey != "" {
		gpgInput.SetValue(account.SigningKey)
	}

	return gpgSetupModel{
		account:    account,
		gpgInput:   gpgInput,
		focusIndex: 0,
	}
}

func (m gpgSetupModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m gpgSetupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		m.gpgInput, cmd = m.gpgInput.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m *gpgSetupModel) updateFocus() {
	if m.focusIndex == 0 {
		m.gpgInput.Focus()
		m.gpgInput.PromptStyle = theme.GetSuccessAnsiStyle()
		m.gpgInput.TextStyle = theme.GetAnsiStyle(theme.White)
	} else {
		m.gpgInput.Blur()
		m.gpgInput.PromptStyle = theme.GetAnsiStyle(theme.BrightBlack)
		m.gpgInput.TextStyle = theme.GetAnsiStyle(theme.BrightBlack)
	}
}

func (m gpgSetupModel) View() string {
	var b strings.Builder

	titleStyle := theme.GetWarningAnsiStyle().Bold(true)
	dimStyle := theme.GetAnsiStyle(theme.BrightBlack)
	highlightStyle := theme.GetWarningAnsiStyle()
	infoStyle := theme.GetAnsiStyle(theme.Cyan)

	b.WriteString(titleStyle.Render(fmt.Sprintf("GPG Signing Key for %s (%s)", m.account.User, m.account.Platform)))
	b.WriteString("\n\n")

	// Instructions
	b.WriteString(infoStyle.Render("To list your GPG keys, run:"))
	b.WriteString("\n")
	b.WriteString("  gpg --list-secret-keys --keyid-format=long\n\n")

	b.WriteString(dimStyle.Render("Enter the key ID (e.g., 3AA5C34371567BD2)"))
	b.WriteString("\n\n")

	// GPG input
	if m.focusIndex == 0 {
		b.WriteString(highlightStyle.Render(">") + " ")
	} else {
		b.WriteString("  ")
	}
	b.WriteString(m.gpgInput.View())
	b.WriteString("\n\n")

	// Buttons
	skipButton := components.NewSkipButton(m.focusIndex == 1)
	saveButton := components.NewSaveButton(m.focusIndex == 2)
	b.WriteString(fmt.Sprintf("%s  %s\n", skipButton.Render(), saveButton.Render()))

	b.WriteString("\n")
	b.WriteString(dimStyle.Render("Press "))
	b.WriteString(highlightStyle.Render("Ctrl+Q"))
	b.WriteString(dimStyle.Render(" to quit"))

	return b.String()
}

func HandleGPGSetup(account *config.Account) error {
	model := newGPGSetupModel(account)
	program := tea.NewProgram(model)
	m, err := program.Run()
	if err != nil {
		return log.Error("failed to run GPG setup", err)
	}

	result, ok := m.(gpgSetupModel)
	if !ok {
		return nil
	}

	if result.skipped {
		log.Info("GPG setup skipped")
		return nil
	}

	if result.saved {
		gpgKey := result.gpgInput.Value()
		if gpgKey != "" {
			account.SigningKey = gpgKey
			log.Success("GPG signing key saved successfully!")
		} else {
			log.Info("No GPG key provided")
		}
	}

	return nil
}
