package createconfig

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestSelectTheme_JKNavigation(t *testing.T) {
	model := newSelectThemeModel()

	updated, _ := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})
	model = updated.(selectThemeModel)
	if model.cursor != 1 {
		t.Fatalf("expected cursor to move to 1, got %d", model.cursor)
	}

	updated, _ = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'k'}})
	model = updated.(selectThemeModel)
	if model.cursor != 0 {
		t.Fatalf("expected cursor to move back to 0, got %d", model.cursor)
	}
}

func TestSelectTheme_QQuits(t *testing.T) {
	model := newSelectThemeModel()

	_, cmd := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	if cmd == nil {
		t.Fatal("expected quit command")
	}
}
