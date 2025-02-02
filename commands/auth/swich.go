package auth

import (
	"fmt"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/config"
	"github.com/spf13/cobra"
)

var swichCmd = &cobra.Command{
	Use:     "swich",
	Aliases: []string{"sw"},
	Short:   "swich accounts",
}

func NewAuthSwich() *cobra.Command {
	swichCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			err := executeCommand(args[0])
			if err != nil {
				return fmt.Errorf("error swiching accounts %w", err)
			}
		} else {
			err := startUserSwich()
			if err != nil {
				return fmt.Errorf("error swiching accounts %w", err)
			}
		}
		return nil
	}

	return swichCmd
}

type model struct {
	allChoices []string
	choices    []string
	cursor     int
	query      string
	selected   string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.choices[m.cursor]
			return m, tea.Quit
		case "backspace":
			if len(m.query) > 0 {
				m.query = m.query[:len(m.query)-1]
			}
		default:
			if msg.Type == tea.KeyRunes {
				m.query += msg.String()
			}
		}

		// Filter accounts based on query
		m.filterChoices()
	}

	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintf("Search: %s\n\n", m.query)

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		line := fmt.Sprintf("%s %s", cursor, choice)

		if m.cursor == i {
			line = fmt.Sprintf("\033[36m%s\033[0m", line)
		}

		s += line + "\n"
	}

	s += "\nPress q to quit.\n"

	return s
}

func (m *model) filterChoices() {
	if m.query == "" {
		m.choices = m.allChoices
		return
	}
	var filtered []string
	for _, choice := range m.allChoices {
		if strings.Contains(choice, m.query) {
			filtered = append(filtered, choice)
		}
	}
	m.choices = filtered
	if m.cursor >= len(filtered) {
		m.cursor = len(filtered) - 1
	}
}

func startUserSwich() error {
	acocunts := config.GlobalConfig.GitHub.Accounts
	var users []string
	for _, acc := range acocunts {
		users = append(users, acc.User)
	}

	initialModel := model{
		allChoices: users,
		choices:    users,
		cursor:     0,
		query:      "",
	}

	p := tea.NewProgram(initialModel)

	if finalModel, err := p.Run(); err == nil {
		if m, ok := finalModel.(model); ok && m.selected != "" {
			fmt.Println("Swiching accounts...")
			err := executeCommand(m.selected)
			if err != nil {
				return err
			}
			fmt.Printf("Swiched to %s\n", m.selected)
		}
	} else {
		return err
	}
	return nil
}

func executeCommand(user string) error {
	var token string
	acocunts := config.GlobalConfig.GitHub.Accounts
	for _, acc := range acocunts {
		if user == acc.User {
			token = acc.Token
			break
		}
	}

	exeCmd := exec.Command("gh", "auth", "login", "--with-token")
	exeCmd.Stdin = strings.NewReader(token)
	if err := exeCmd.Run(); err != nil {
		return fmt.Errorf("error checking swich account with err (%w) ", err)
	}
	return nil
}
