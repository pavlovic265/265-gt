package account

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pavlovic265/265-gt/config"
)

func TestTokenSetup_QQuitsOnlyInNormalMode(t *testing.T) {
	model := newTokenSetupModel(&config.Account{})

	updated, cmd := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	next := updated.(tokenSetupModel)
	if cmd == nil {
		t.Fatal("expected quit command in normal mode")
	}
	if !next.skipped {
		t.Fatal("expected setup to be marked as skipped")
	}

	model = newTokenSetupModel(&config.Account{})
	updated, _ = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'i'}})
	model = updated.(tokenSetupModel)

	updated, cmd = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	next = updated.(tokenSetupModel)
	if next.tokenInput.Value() != "q" {
		t.Fatalf("expected token input to contain q, got %q", next.tokenInput.Value())
	}
	if next.skipped {
		t.Fatal("did not expect setup to be skipped in insert mode")
	}
	if cmd == nil {
		t.Fatal("expected input update command in insert mode")
	}
}

func TestTokenSetup_EscLeavesInsertMode(t *testing.T) {
	model := newTokenSetupModel(&config.Account{})

	updated, _ := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'i'}})
	model = updated.(tokenSetupModel)
	if !model.insertMode {
		t.Fatal("expected insert mode to be enabled")
	}

	updated, _ = model.Update(tea.KeyMsg{Type: tea.KeyEsc})
	model = updated.(tokenSetupModel)
	if model.insertMode {
		t.Fatal("expected insert mode to be disabled")
	}
}

func TestTokenSetup_JKMovesFocusInNormalMode(t *testing.T) {
	model := newTokenSetupModel(&config.Account{})

	updated, _ := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})
	model = updated.(tokenSetupModel)
	if model.focusIndex != 1 {
		t.Fatalf("expected focus to move to 1, got %d", model.focusIndex)
	}

	updated, _ = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'k'}})
	model = updated.(tokenSetupModel)
	if model.focusIndex != 0 {
		t.Fatalf("expected focus to move back to 0, got %d", model.focusIndex)
	}
}
