package account

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/helpers"
	"github.com/pavlovic265/265-gt/runner"
	"github.com/pavlovic265/265-gt/utils/log"
)

const (
	stepChooseOption = iota
	stepSelectKeyType
	stepEnterKeyPath
)

type sshSetupModel struct {
	account    *config.Account
	step       int
	createNew  bool
	keyType    helpers.SSHKeyType
	keyTypeIdx int
	keyPath    textinput.Model
	quitting   bool
	sshHelper  helpers.SSHHelper
	err        string
}

var keyTypeOptions = []helpers.SSHKeyType{
	helpers.SSHKeyTypeEd25519,
	helpers.SSHKeyTypeRSA,
}

func newSSHSetupModel(account *config.Account, sshHelper helpers.SSHHelper) sshSetupModel {
	keyPathInput := textinput.New()
	keyPathInput.Placeholder = helpers.DefaultSSHKeyPath(strings.ToLower(string(account.Platform)), account.User)
	keyPathInput.Focus()

	return sshSetupModel{
		account:    account,
		step:       stepChooseOption,
		keyType:    helpers.SSHKeyTypeEd25519,
		keyTypeIdx: 0,
		keyPath:    keyPathInput,
		sshHelper:  sshHelper,
	}
}

func (m sshSetupModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m sshSetupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String(), tea.KeyCtrlQ.String(), tea.KeyEsc.String():
			m.quitting = true
			return m, tea.Quit

		case "1":
			if m.step == stepChooseOption {
				m.createNew = true
				m.step = stepSelectKeyType
				return m, nil
			}
		case "2":
			if m.step == stepChooseOption {
				m.createNew = false
				m.step = stepEnterKeyPath
				return m, nil
			}
		case "3":
			if m.step == stepChooseOption {
				m.quitting = true
				return m, tea.Quit
			}

		case tea.KeyUp.String(), tea.KeyDown.String(), "j", "k":
			if m.step == stepSelectKeyType {
				m.keyTypeIdx = (m.keyTypeIdx + 1) % len(keyTypeOptions)
				m.keyType = keyTypeOptions[m.keyTypeIdx]
				return m, nil
			}

		case tea.KeyEnter.String():
			if m.step == stepSelectKeyType {
				m.step = stepEnterKeyPath
				m.keyPath.Focus()
				return m, textinput.Blink
			}
			if m.step == stepEnterKeyPath {
				return m, tea.Quit
			}
		}
	}

	if m.step == stepEnterKeyPath {
		var cmd tea.Cmd
		m.keyPath, cmd = m.keyPath.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m sshSetupModel) View() string {
	var b strings.Builder

	titleStyle := constants.GetWarningAnsiStyle().Bold(true)
	dimStyle := constants.GetAnsiStyle(constants.BrightBlack)
	highlightStyle := constants.GetWarningAnsiStyle()
	errStyle := constants.GetErrorAnsiStyle()

	b.WriteString(titleStyle.Render(fmt.Sprintf("SSH Setup for %s", m.account.User)))
	b.WriteString("\n\n")

	switch m.step {
	case stepChooseOption:
		b.WriteString("Choose SSH key option:\n\n")
		b.WriteString(highlightStyle.Render("  [1]") + " Create new SSH key\n")
		b.WriteString(highlightStyle.Render("  [2]") + " Use existing SSH key\n")
		b.WriteString(highlightStyle.Render("  [3]") + " Skip SSH setup\n")

	case stepSelectKeyType:
		b.WriteString("Select key type:\n\n")
		for i, kt := range keyTypeOptions {
			if i == m.keyTypeIdx {
				b.WriteString(highlightStyle.Render("  (•) "))
				b.WriteString(string(kt))
				if kt == helpers.SSHKeyTypeEd25519 {
					b.WriteString(dimStyle.Render(" (recommended)"))
				}
			} else {
				b.WriteString(dimStyle.Render("  ( ) " + string(kt)))
			}
			b.WriteString("\n")
		}
		b.WriteString("\n")
		b.WriteString(dimStyle.Render("Press Enter to continue"))

	case stepEnterKeyPath:
		if m.createNew {
			b.WriteString("Key type: ")
			b.WriteString(highlightStyle.Render(string(m.keyType)))
			b.WriteString("\n\n")
		}
		b.WriteString("SSH key path:\n")
		b.WriteString(m.keyPath.View())
		b.WriteString("\n\n")
		b.WriteString(dimStyle.Render("Press Enter to confirm"))
	}

	if m.err != "" {
		b.WriteString("\n\n")
		b.WriteString(errStyle.Render("Error: " + m.err))
	}

	b.WriteString("\n\n")
	b.WriteString(dimStyle.Render("Press "))
	b.WriteString(highlightStyle.Render("Ctrl+Q"))
	b.WriteString(dimStyle.Render(" to quit"))

	return b.String()
}

func HandleSSHSetup(account *config.Account, r runner.Runner) error {
	sshHelper := helpers.NewSSHHelper(r)

	model := newSSHSetupModel(account, sshHelper)
	program := tea.NewProgram(model)
	m, err := program.Run()
	if err != nil {
		return log.Error("failed to run SSH setup", err)
	}

	result, ok := m.(sshSetupModel)
	if !ok || result.quitting {
		return nil
	}

	keyPath := result.keyPath.Value()
	if keyPath == "" {
		keyPath = helpers.DefaultSSHKeyPath(strings.ToLower(string(account.Platform)), account.User)
	}

	if result.createNew {
		email := account.Email
		if email == "" {
			noReplyMail := constants.GitHubNoReplyMail
			if account.Platform == constants.GitLabPlatform {
				noReplyMail = constants.GitLabNoReplyMail
			}
			email = account.User + noReplyMail
		}

		log.Infof("Generating %s SSH key at %s...", result.keyType, keyPath)
		if err := sshHelper.GenerateKey(keyPath, email, result.keyType); err != nil {
			return log.Error("failed to generate SSH key", err)
		}

		log.Info("Adding key to ssh-agent...")
		if err := sshHelper.AddToAgent(keyPath); err != nil {
			log.Warning("Could not add key to ssh-agent (you may need to add it manually)")
		}
	}

	hostname := constants.GitHubHost
	if account.Platform == constants.GitLabPlatform {
		hostname = constants.GitLabHost
	}

	sshHost := helpers.BuildSSHHost(hostname, account.User)

	log.Infof("Adding SSH config for %s...", sshHost)
	if err := sshHelper.AddHostConfig(sshHost, hostname, keyPath); err != nil {
		return log.Error("failed to add SSH config", err)
	}

	account.SSHKeyPath = keyPath
	account.SSHHost = sshHost

	pubKey, err := sshHelper.GetPublicKey(keyPath)
	if err == nil && result.createNew {
		log.Success("SSH key generated successfully!")
		log.Infof("Add this public key to your %s account:", account.Platform)
		log.Infof("\n%s\n", pubKey)
	} else {
		log.Success("SSH configuration updated!")
	}

	return nil
}
