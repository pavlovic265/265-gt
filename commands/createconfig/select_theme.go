package createconfig

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/utils/log"
)

type selectThemeModel struct {
	cursor  int
	choice  constants.Theme
	choices []constants.Theme
}

var themeChoices = []constants.Theme{
	constants.DarkTheme,
	constants.LightTheme,
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
		s.WriteString(sm.choices[i].String())
		s.WriteString("\n")
	}
	s.WriteString("\n(press ctrl+q to quit)\n")

	return s.String()
}

func HandleSelectTheme() (*config.ThemeConfig, error) {
	themeModel := newSelectThemeModel()
	selectProgram := tea.NewProgram(themeModel)
	m, err := selectProgram.Run()
	if err != nil {
		return nil, err
	}

	if m, ok := m.(selectThemeModel); ok && m.choice != "" {
		themeConfig := config.ThemeConfig{Type: m.choice}
		return &themeConfig, nil
	}

	return nil, log.ErrorMsg("failed to select theme")
}
