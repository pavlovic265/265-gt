package components

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/ui/theme"
)

type ListModel[T any] struct {
	AllChoices    []T
	Choices       []T
	Cursor        int
	Query         string
	SearchMode    bool
	Selected      T
	YankAction    bool
	MergeAction   bool
	RefreshAction bool
	EnableYank    bool
	EnableMerge   bool
	EnableRefresh bool
	Refreshing    bool
	// Formatter is a function that converts T to string for display
	Formatter func(T) string
	// Matcher is a function that checks if T matches the query string
	Matcher func(T, string) bool
	// RefreshFunc is called when refresh is triggered
	RefreshFunc func() tea.Msg

	windowHeight int // terminal height from tea.WindowSizeMsg
	offset       int // index of first visible item in the viewport
}

type RefreshCompleteMsg[T any] struct {
	Choices []T
	Err     error
}

var (
	searchLabelStyle = lipgloss.NewStyle().
				Foreground(theme.Blue)

	cursorStyle = lipgloss.NewStyle().
			Foreground(theme.Yellow)

	itemStyle = lipgloss.NewStyle().
			Foreground(theme.Foreground)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(theme.Green)

	emptyStateStyle = lipgloss.NewStyle().
			Foreground(theme.BrightBlack)

	footerStyle = lipgloss.NewStyle().
			Foreground(theme.BrightBlack)

	keyStyle = lipgloss.NewStyle().
			Foreground(theme.Yellow)
)

func (m ListModel[T]) Init() tea.Cmd {
	return nil
}

func (m ListModel[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowHeight = msg.Height
		m.adjustViewport()
		return m, nil

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
			m.adjustViewport()
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String():
			return m, tea.Quit
		case tea.KeyEsc.String():
			if m.SearchMode {
				m.SearchMode = false
			}
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
		case tea.KeyBackspace.String():
			if m.SearchMode && len(m.Query) > 0 {
				m.Query = m.Query[:len(m.Query)-1]
			}
		default:
			if msg.Type == tea.KeyRunes {
				switch {
				case m.SearchMode:
					m.Query += msg.String()
				case msg.String() == "/":
					m.SearchMode = true
				case msg.String() == "q":
					return m, tea.Quit
				case msg.String() == "y":
					if m.EnableYank && len(m.Choices) > 0 && m.Cursor >= 0 && m.Cursor < len(m.Choices) {
						m.Selected = m.Choices[m.Cursor]
						m.YankAction = true
						return m, tea.Quit
					}
				case msg.String() == "r":
					if m.EnableRefresh && !m.Refreshing && m.RefreshFunc != nil {
						m.Refreshing = true
						return m, m.RefreshFunc
					}
				case msg.String() == "m":
					if m.EnableMerge && len(m.Choices) > 0 && m.Cursor >= 0 && m.Cursor < len(m.Choices) {
						m.Selected = m.Choices[m.Cursor]
						m.MergeAction = true
						return m, tea.Quit
					}
				}
			}
		}
		m.filterChoices()
		m.adjustViewport()
	}

	return m, nil
}

func (m ListModel[T]) View() string {
	var content strings.Builder

	// Search input
	searchLabel := searchLabelStyle.Render("Search:")
	searchValue := m.Query
	if !m.SearchMode && m.Query == "" {
		searchValue = "(/ to search)"
	}
	content.WriteString(fmt.Sprintf("%s %s\n\n", searchLabel, searchValue))

	// List items
	if len(m.Choices) == 0 {
		content.WriteString(emptyStateStyle.Render("No items found"))
		content.WriteString("\n")
	} else {
		visibleCount := len(m.Choices) // default: show all
		if m.windowHeight > 5 {
			visibleCount = m.windowHeight - 5
		}
		end := m.offset + visibleCount
		if end > len(m.Choices) {
			end = len(m.Choices)
		}
		for i := m.offset; i < end; i++ {
			choice := m.Choices[i]
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
		refreshingStyle := lipgloss.NewStyle().Foreground(theme.Yellow)
		content.WriteString(refreshingStyle.Render("⟳ Refreshing..."))
		content.WriteString("\n\n")
	}

	if m.SearchMode {
		content.WriteString(footerStyle.Render("Press "))
		content.WriteString(keyStyle.Render("ESC"))
		content.WriteString(footerStyle.Render(" to exit search mode"))
	} else {
		content.WriteString(footerStyle.Render("Press "))
		content.WriteString(keyStyle.Render("/"))
		content.WriteString(footerStyle.Render(" to search, "))
		content.WriteString(keyStyle.Render("q"))
		content.WriteString(footerStyle.Render(" to quit"))

		if m.EnableRefresh && len(m.Choices) > 0 {
			content.WriteString(footerStyle.Render(", "))
			content.WriteString(keyStyle.Render("r"))
			content.WriteString(footerStyle.Render(" to refresh"))
		}

		if m.EnableYank && len(m.Choices) > 0 {
			content.WriteString(footerStyle.Render(", "))
			content.WriteString(keyStyle.Render("y"))
			content.WriteString(footerStyle.Render(" to yank URL"))
		}

		if m.EnableMerge && len(m.Choices) > 0 {
			content.WriteString(footerStyle.Render(", "))
			content.WriteString(keyStyle.Render("m"))
			content.WriteString(footerStyle.Render(" to merge"))
		}
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
		Foreground(theme.Yellow).
		Background(theme.BrightBlack)

	return before + highlightStyle.Render(match) + after
}

func (m *ListModel[T]) adjustViewport() {
	visibleCount := m.windowHeight - 5
	if visibleCount <= 0 || m.windowHeight == 0 {
		return
	}
	if m.Cursor < m.offset {
		m.offset = m.Cursor
	}
	if m.Cursor >= m.offset+visibleCount {
		m.offset = m.Cursor - visibleCount + 1
	}
	maxOffset := len(m.Choices) - visibleCount
	if maxOffset < 0 {
		maxOffset = 0
	}
	if m.offset > maxOffset {
		m.offset = maxOffset
	}
	if m.offset < 0 {
		m.offset = 0
	}
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
