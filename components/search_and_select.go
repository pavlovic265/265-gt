package components

import (
	"fmt"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type ListModel struct {
	AllChoices []string
	Choices    []string
	Cursor     int
	Query      string
	Selected   string
	YankURL    string
	URLs       []string
	Yanked     bool
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyEsc.String(), tea.KeyCtrlC.String(), tea.KeyCtrlQ.String():
			return m, tea.Quit
		case tea.KeyShiftTab.String(), tea.KeyUp.String(), tea.KeyCtrlK.String():
			if m.Cursor > 0 {
				m.Cursor--
			} else {
				m.Cursor = len(m.Choices) - 1
			}
			m.updateYankURL()
		case tea.KeyTab.String(), tea.KeyDown.String(), tea.KeyCtrlJ.String():
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			} else {
				m.Cursor = 0
			}
			m.updateYankURL()
		case tea.KeyEnter.String():
			m.Selected = m.Choices[m.Cursor]
			return m, tea.Quit
		case tea.KeyCtrlY.String():
			if len(m.URLs) == 0 && m.YankURL == "" {
				return m, nil
			}
			m.yankToClipboard(m.YankURL)
			m.Yanked = true
			return m, tea.Quit
		case tea.KeyBackspace.String():
			if len(m.Query) > 0 {
				m.Query = m.Query[:len(m.Query)-1]
			}
		default:
			if msg.Type == tea.KeyRunes {
				m.Query += msg.String()
			}
		}
		m.filterChoices()
	}

	return m, nil
}

func (m ListModel) View() string {
	s := fmt.Sprintf("Search: %s\n\n", m.Query)

	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		line := fmt.Sprintf("%s %s", cursor, choice)

		if m.Cursor == i {
			line = fmt.Sprintf("\033[36m%s\033[0m", line)
		}

		s += line + "\n"
	}

	// Only show yank option if there are URLs available
	if len(m.URLs) > 0 && m.YankURL != "" {
		s += "\nPress CTRL+q to quit, CTRL+y to yank URL.\n"
	} else {
		s += "\nPress CTRL+q to quit.\n"
	}

	return s
}

func (m *ListModel) filterChoices() {
	if m.Query == "" {
		m.Choices = m.AllChoices
		return
	}

	var filtered []string
	var filteredURLs []string
	for i, choice := range m.AllChoices {
		if strings.Contains(choice, m.Query) {
			filtered = append(filtered, choice)
			if i < len(m.URLs) {
				filteredURLs = append(filteredURLs, m.URLs[i])
			}
		}
	}
	m.Choices = filtered
	m.URLs = filteredURLs
	if m.Cursor >= len(filtered) {
		m.Cursor = len(filtered) - 1
	}
	m.updateYankURL()
}

func (m *ListModel) yankToClipboard(url string) {
	commands := [][]string{
		{"pbcopy"},                           // macOS
		{"xclip", "-selection", "clipboard"}, // Linux with xclip
		{"xsel", "--clipboard", "--input"},   // Linux with xsel
		{"clip"},                             // Windows
	}

	for _, cmd := range commands {
		command := exec.Command(cmd[0], cmd[1:]...)
		command.Stdin = strings.NewReader(url)
		if err := command.Run(); err == nil {
			return
		}
	}
}

func (m *ListModel) updateYankURL() {
	if m.Cursor >= 0 && m.Cursor < len(m.Choices) && m.Cursor < len(m.URLs) {
		m.YankURL = m.URLs[m.Cursor]
	}
}
