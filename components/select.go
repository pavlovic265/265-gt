package components

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/constants"
)

type ListModel[T any] struct {
	AllChoices    []T
	Choices       []T
	Cursor        int
	Query         string
	Selected      T
	YankAction    bool
	MergeAction   bool
	UpdateAction  bool
	RefreshAction bool
	EnableYank    bool
	EnableMerge   bool
	EnableUpdate  bool
	EnableRefresh bool
	Refreshing    bool
	// Formatter is a function that converts T to string for display
	Formatter func(T) string
	// Matcher is a function that checks if T matches the query string
	Matcher func(T, string) bool
	// RefreshFunc is called when refresh is triggered
	RefreshFunc func() tea.Msg
}

// RefreshCompleteMsg is sent when refresh completes
type RefreshCompleteMsg[T any] struct {
	Choices []T
	Err     error
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

func (m ListModel[T]) Init() tea.Cmd {
	return nil
}

func (m ListModel[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case RefreshCompleteMsg[T]:
		m.Refreshing = false
		if msg.Err == nil {
			// Store current selected item's index to preserve cursor position
			var selectedIndex int
			if len(m.Choices) > 0 && m.Cursor >= 0 && m.Cursor < len(m.Choices) {
				selectedItem := m.Choices[m.Cursor]
				// Find the same item in the new choices
				selectedIndex = -1
				for i, choice := range msg.Choices {
					if m.Formatter(choice) == m.Formatter(selectedItem) {
						selectedIndex = i
						break
					}
				}
			}

			m.AllChoices = msg.Choices
			m.filterChoices()

			// Restore cursor position if we found the item
			if selectedIndex >= 0 && selectedIndex < len(m.Choices) {
				m.Cursor = selectedIndex
			} else if m.Cursor >= len(m.Choices) && len(m.Choices) > 0 {
				m.Cursor = len(m.Choices) - 1
			}
		}
		return m, nil

	case tea.KeyMsg:
		// fmt.Printf("DBG key: type=%v str=%q runes=%q alt=%v\n", msg.Type, msg.String(), string(msg.Runes), msg.Alt)

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
			}
		case tea.KeyTab.String(), tea.KeyDown.String(), tea.KeyCtrlJ.String():
			if len(m.Choices) > 0 {
				if m.Cursor < len(m.Choices)-1 {
					m.Cursor++
				} else {
					m.Cursor = 0
				}
			}
		case tea.KeyEnter.String():
			if len(m.Choices) > 0 && m.Cursor >= 0 && m.Cursor < len(m.Choices) {
				m.Selected = m.Choices[m.Cursor]
			}
			return m, tea.Quit
		case tea.KeyCtrlY.String():
			if len(m.Choices) > 0 && m.Cursor >= 0 && m.Cursor < len(m.Choices) {
				m.Selected = m.Choices[m.Cursor]
			}
			m.YankAction = true
			return m, tea.Quit
		case tea.KeyCtrlR.String():
			if m.EnableRefresh && !m.Refreshing && m.RefreshFunc != nil {
				m.Refreshing = true
				return m, m.RefreshFunc
			}
		case tea.KeyCtrlO.String():
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

func (m ListModel[T]) View() string {
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

			// Convert choice to string using formatter
			choiceStr := m.Formatter(choice)

			// Highlight search matches
			highlightedChoice := highlightMatch(choiceStr, m.Query)

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

	// Show refreshing indicator
	if m.Refreshing {
		refreshingStyle := lipgloss.NewStyle().Foreground(constants.Yellow)
		content.WriteString(refreshingStyle.Render("⟳ Refreshing..."))
		content.WriteString("\n\n")
	}

	content.WriteString(footerStyle.Render("Press "))
	content.WriteString(keyStyle.Render("CTRL+Q"))
	content.WriteString(footerStyle.Render(" to quit"))

	// Only show refresh option if enabled
	if m.EnableRefresh && len(m.Choices) > 0 {
		content.WriteString(footerStyle.Render(", "))
		content.WriteString(keyStyle.Render("CTRL+R"))
		content.WriteString(footerStyle.Render(" to refresh"))
	}

	// Only show yank option if enabled
	if m.EnableYank && len(m.Choices) > 0 {
		content.WriteString(footerStyle.Render(", "))
		content.WriteString(keyStyle.Render("CTRL+Y"))
		content.WriteString(footerStyle.Render(" to yank URL"))
	}

	// Show merge option if enabled
	if m.EnableMerge && len(m.Choices) > 0 {
		content.WriteString(footerStyle.Render(", "))
		content.WriteString(keyStyle.Render("CTRL+O"))
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

func (m *ListModel[T]) filterChoices() {
	if m.Query == "" {
		m.Choices = m.AllChoices
		return
	}

	var filtered []T
	for _, choice := range m.AllChoices {
		if m.Matcher(choice, m.Query) {
			filtered = append(filtered, choice)
		}
	}
	m.Choices = filtered
	if len(filtered) == 0 {
		m.Cursor = 0
	} else if m.Cursor >= len(filtered) {
		m.Cursor = len(filtered) - 1
	}
}
