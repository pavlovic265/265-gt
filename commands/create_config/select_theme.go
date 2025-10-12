package createconfig

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/constants"
)

type selectThemeModel struct {
	cursor  int
	choice  string
	choices []string
}

var themeChoices = []string{
	constants.DarkTheme.String(),
	constants.LightTheme.String(),
}

func newSelectThemeModel() selectThemeModel {
	return selectThemeModel{
		choices: themeChoices,
	}
}

func (sm selectThemeModel) Init() tea.Cmd {
	return nil
}

func (sm selectThemeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String(), tea.KeyCtrlQ.String(), tea.KeyEsc.String():
			return sm, tea.Quit

		case tea.KeyEnter.String():
			sm.choice = sm.choices[sm.cursor]
			return sm, tea.Quit

		case tea.KeyDown.String(), tea.KeyCtrlJ.String(), tea.KeyTab.String():
			sm.cursor++
			if sm.cursor >= len(sm.choices) {
				sm.cursor = 0
			}

		case tea.KeyUp.String(), tea.KeyCtrlK.String(), tea.KeyShiftTab.String():
			sm.cursor--
			if sm.cursor < 0 {
				sm.cursor = len(sm.choices) - 1
			}
		}
	}

	return sm, nil
}

func (sm selectThemeModel) View() string {
	s := strings.Builder{}
	s.WriteString("Choose theme?" + "\n\n")

	for i := 0; i < len(sm.choices); i++ {
		if sm.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		if sm.cursor == i {
			s.WriteString(sm.choices[i])
		} else {
			s.WriteString(sm.choices[i])
		}
		s.WriteString("\n")
	}
	s.WriteString("\n(press ctrl+q to quit)\n")

	return s.String()
}

func HandleSelectTheme() (*constants.Theme, error) {
	themeModel := newSelectThemeModel()
	selectProgram := tea.NewProgram(themeModel)
	m, err := selectProgram.Run()
	if err != nil {
		return nil, err
	}

	if m, ok := m.(selectThemeModel); ok && m.choice != "" {
		// Convert string choice back to Theme type
		var theme constants.Theme
		switch m.choice {
		case constants.DarkTheme.String():
			theme = constants.DarkTheme
		case constants.LightTheme.String():
			theme = constants.LightTheme
		default:
			return nil, fmt.Errorf("invalid theme choice: %s", m.choice)
		}
		return &theme, nil
	}

	return nil, fmt.Errorf("failed to select theme")
}
