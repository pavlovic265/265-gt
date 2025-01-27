package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:     "checkout",
	Aliases: []string{"co"},
	Short:   "checkout branch",
}

func Checkout() *cobra.Command {
	checkoutCmd.RunE = func(cmd *cobra.Command, args []string) error {
		checkoutBranch()
		return nil
	}

	return checkoutCmd
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

		// Filter branches based on query
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

func getBranches() ([]string, error) {
	out, err := exec.Command("git", "branch", "--list").Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")
	var branches []string
	for _, line := range lines {
		branch := strings.TrimSpace(line)
		if branch != "" {
			branches = append(branches, strings.TrimPrefix(branch, "* "))
		}
	}
	return branches, nil
}

func checkoutBranch() {
	branches, err := getBranches()
	if err != nil {
		fmt.Println("Error getting branches:", err)
		return
	}

	initialModel := model{
		allChoices: branches,
		choices:    branches,
		cursor:     0,
		query:      "",
	}

	p := tea.NewProgram(initialModel)

	if finalModel, err := p.Run(); err == nil {
		if m, ok := finalModel.(model); ok && m.selected != "" {
			fmt.Printf("Checking out branch '%s'...\n", m.selected)
			err = exec.Command("git", "checkout", m.selected).Run()
			if err != nil {
				fmt.Println("Error checking out branch:", err)
			}
		}
	} else {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
