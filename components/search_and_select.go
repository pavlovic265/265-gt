package components

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type ListModel struct {
	AllChoices []string
	Choices    []string
	Cursor     int
	Query      string
	Selected   string
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
		case tea.KeyTab.String(), tea.KeyUp.String(), tea.KeyCtrlK.String():
			if m.Cursor > 0 {
				m.Cursor--
			} else {
				m.Cursor = len(m.Choices) - 1
			}
		case tea.KeyShiftTab.String(), tea.KeyDown.String(), tea.KeyCtrlJ.String():
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			} else {
				m.Cursor = 0
			}
		case tea.KeyEnter.String():
			m.Selected = m.Choices[m.Cursor]
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
		// Filter branches based on query
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

	s += "\nPress CTRL+q to quit.\n"

	return s
}

func (m *ListModel) filterChoices() {
	if m.Query == "" {
		m.Choices = m.AllChoices
		return
	}

	var filtered []string
	for _, choice := range m.AllChoices {
		if strings.Contains(choice, m.Query) {
			filtered = append(filtered, choice)
		}
	}
	m.Choices = filtered
	if m.Cursor >= len(filtered) {
		m.Cursor = len(filtered) - 1
	}
}
