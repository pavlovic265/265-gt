package createconfig

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/utils/log"
)

type selectPlatformModel struct {
	cursor  int
	choice  string
	choices []string
}

var choices = []string{
	constants.GitHubPlatform.String(),
	constants.GitLabPlatform.String() + " (not implemented yet)",
}

func newSelectPlatformModel() selectPlatformModel {
	return selectPlatformModel{
		choices: choices,
	}
}

func (sm selectPlatformModel) Init() tea.Cmd {
	return nil
}

func (sm selectPlatformModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String(), tea.KeyCtrlQ.String(), tea.KeyCtrlQ.String(), tea.KeyEsc.String():
			return sm, tea.Quit

		case tea.KeyEnter.String():
			// Send the choice on the channel and exit.
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

func (sm selectPlatformModel) View() string {
	s := strings.Builder{}
	s.WriteString("Choose platform?" + "\n\n")

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

func HandleSelectPlatform() (*string, error) {
	platformModel := newSelectPlatformModel()
	selectProgram := tea.NewProgram(platformModel)
	m, err := selectProgram.Run()
	if err != nil {
		return nil, err
	}

	if m, ok := m.(selectPlatformModel); ok && m.choice != "" {
		return &m.choice, nil
	}

	return nil, log.ErrorMsg("failed to select platform")
}
