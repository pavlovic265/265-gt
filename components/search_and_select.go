package components

import (
	"fmt"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/constants"
)

type ListModel struct {
	AllChoices   []string
	Choices      []string
	Cursor       int
	Query        string
	Selected     string
	YankURL      string
	URLs         []string
	Yanked       bool
	MergeAction  bool
	UpdateAction bool
	EnableMerge  bool
	EnableUpdate bool
}

// Styling definitions
var (
	// Search input styles
	searchLabelStyle = lipgloss.NewStyle().
				Foreground(constants.Blue)

	// List item styles
	cursorStyle = lipgloss.NewStyle().
			Foreground(constants.Yellow)

	itemStyle = lipgloss.NewStyle().
			Foreground(constants.Foreground)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(constants.Green)

	// Empty state style
	emptyStateStyle = lipgloss.NewStyle().
			Foreground(constants.BrightBlack)

	// Footer styles
	footerStyle = lipgloss.NewStyle().
			Foreground(constants.BrightBlack)

	keyStyle = lipgloss.NewStyle().
			Foreground(constants.Yellow)
)

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
			if len(m.Choices) > 0 {
				if m.Cursor > 0 {
					m.Cursor--
				} else {
					m.Cursor = len(m.Choices) - 1
				}
				m.updateYankURL()
			}
		case tea.KeyTab.String(), tea.KeyDown.String(), tea.KeyCtrlJ.String():
			if len(m.Choices) > 0 {
				if m.Cursor < len(m.Choices)-1 {
					m.Cursor++
				} else {
					m.Cursor = 0
				}
				m.updateYankURL()
			}
		case tea.KeyEnter.String():
			if len(m.Choices) > 0 && m.Cursor >= 0 && m.Cursor < len(m.Choices) {
				m.Selected = m.Choices[m.Cursor]
			}
			return m, tea.Quit
		case tea.KeyCtrlY.String():
			if len(m.URLs) == 0 && m.YankURL == "" {
				return m, nil
			}
			m.yankToClipboard(m.YankURL)
			m.Yanked = true
			return m, tea.Quit
		case tea.KeyCtrlM.String():
			if m.EnableMerge && len(m.Choices) > 0 && m.Cursor >= 0 && m.Cursor < len(m.Choices) {
				m.Selected = m.Choices[m.Cursor]
				m.MergeAction = true
				return m, tea.Quit
			}
		case tea.KeyCtrlU.String():
			if m.EnableUpdate && len(m.Choices) > 0 && m.Cursor >= 0 && m.Cursor < len(m.Choices) {
				m.Selected = m.Choices[m.Cursor]
				m.UpdateAction = true
				return m, tea.Quit
			}
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
	var content strings.Builder

	// Search input
	searchLabel := searchLabelStyle.Render("Search:")
	content.WriteString(fmt.Sprintf("%s %s\n\n", searchLabel, m.Query))

	// List items
	if len(m.Choices) == 0 {
		content.WriteString(emptyStateStyle.Render("No items found"))
		content.WriteString("\n")
	} else {
		for i, choice := range m.Choices {
			var cursor, styledChoice string

			// Highlight search matches
			highlightedChoice := highlightMatch(choice, m.Query)

			if m.Cursor == i {
				cursor = cursorStyle.Render(">")
				styledChoice = selectedItemStyle.Render(highlightedChoice)
			} else {
				cursor = " "
				styledChoice = itemStyle.Render(highlightedChoice)
			}

			content.WriteString(fmt.Sprintf("%s %s\n", cursor, styledChoice))
		}
	}

	// Footer with instructions
	content.WriteString("\n")

	content.WriteString(footerStyle.Render("Press "))
	content.WriteString(keyStyle.Render("CTRL+Q"))
	content.WriteString(footerStyle.Render(" to quit"))

	// Only show yank option if there are URLs available
	if len(m.URLs) > 0 && m.YankURL != "" {
		content.WriteString(footerStyle.Render(", "))
		content.WriteString(keyStyle.Render("CTRL+Y"))
		content.WriteString(footerStyle.Render(" to yank URL"))
	}

	// Show merge option if enabled
	if m.EnableMerge && len(m.Choices) > 0 {
		content.WriteString(footerStyle.Render(", "))
		content.WriteString(keyStyle.Render("CTRL+M"))
		content.WriteString(footerStyle.Render(" to merge"))
	}

	// Show update option if enabled
	if m.EnableUpdate && len(m.Choices) > 0 {
		content.WriteString(footerStyle.Render(", "))
		content.WriteString(keyStyle.Render("CTRL+U"))
		content.WriteString(footerStyle.Render(" to update"))
	}

	return content.String()
}

// highlightMatch highlights the search query in the text
func highlightMatch(text, query string) string {
	if query == "" {
		return text
	}

	// Find the match and highlight it
	lowerText := strings.ToLower(text)
	lowerQuery := strings.ToLower(query)

	index := strings.Index(lowerText, lowerQuery)
	if index == -1 {
		return text
	}

	// Split the text and highlight the match
	before := text[:index]
	match := text[index : index+len(query)]
	after := text[index+len(query):]

	highlightStyle := lipgloss.NewStyle().
		Foreground(constants.Yellow).
		Background(constants.BrightBlack)

	return before + highlightStyle.Render(match) + after
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
	if len(filtered) == 0 {
		m.Cursor = 0
	} else if m.Cursor >= len(filtered) {
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
	if m.Cursor >= 0 && m.Cursor < len(m.Choices) && m.Cursor < len(m.URLs) && len(m.URLs) > 0 {
		m.YankURL = m.URLs[m.Cursor]
	} else {
		m.YankURL = ""
	}
}
