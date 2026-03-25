package components

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func newTestListModel() ListModel[string] {
	return ListModel[string]{
		AllChoices: []string{"alpha", "beta"},
		Choices:    []string{"alpha", "beta"},
		Formatter:  func(s string) string { return s },
		Matcher: func(s, query string) bool {
			return len(query) == 0 || s == query || s == "alpha" && query == "a" || s == "beta" && query == "b"
		},
	}
}

func TestListModel_SlashEntersSearchMode(t *testing.T) {
	model := newTestListModel()

	updated, _ := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'/'}})
	next := updated.(ListModel[string])

	if !next.SearchMode {
		t.Fatal("expected search mode to be enabled")
	}
	if next.Query != "" {
		t.Fatalf("expected empty query, got %q", next.Query)
	}
}

func TestListModel_TypingInNormalModeDoesNotSearch(t *testing.T) {
	model := newTestListModel()

	updated, _ := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}})
	next := updated.(ListModel[string])

	if next.SearchMode {
		t.Fatal("expected search mode to stay disabled")
	}
	if next.Query != "" {
		t.Fatalf("expected empty query, got %q", next.Query)
	}
}

func TestListModel_TypingInSearchModeUpdatesQuery(t *testing.T) {
	model := newTestListModel()
	updated, _ := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'/'}})
	model = updated.(ListModel[string])

	updated, _ = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}})
	next := updated.(ListModel[string])

	if next.Query != "a" {
		t.Fatalf("expected query to be %q, got %q", "a", next.Query)
	}
}

func TestListModel_EscExitsSearchMode(t *testing.T) {
	model := newTestListModel()
	updated, _ := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'/'}})
	model = updated.(ListModel[string])

	updated, _ = model.Update(tea.KeyMsg{Type: tea.KeyEsc})
	next := updated.(ListModel[string])

	if next.SearchMode {
		t.Fatal("expected search mode to be disabled")
	}
}

func TestListModel_QQuitsOnlyInNormalMode(t *testing.T) {
	model := newTestListModel()

	updated, cmd := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	next := updated.(ListModel[string])
	if cmd == nil {
		t.Fatal("expected quit command in normal mode")
	}
	if next.Query != "" {
		t.Fatalf("expected empty query, got %q", next.Query)
	}

	updated, _ = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'/'}})
	model = updated.(ListModel[string])
	updated, cmd = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	next = updated.(ListModel[string])
	if cmd != nil {
		t.Fatal("did not expect quit command in search mode")
	}
	if next.Query != "q" {
		t.Fatalf("expected query to be %q, got %q", "q", next.Query)
	}
}

func TestListModel_ActionKeysRespectFlags(t *testing.T) {
	model := newTestListModel()
	model.EnableYank = true
	model.EnableRefresh = true
	model.EnableMerge = true
	model.RefreshFunc = func() tea.Msg { return nil }

	updated, cmd := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'y'}})
	next := updated.(ListModel[string])
	if !next.YankAction || cmd == nil {
		t.Fatal("expected yank action to trigger")
	}

	model = newTestListModel()
	model.EnableRefresh = true
	model.RefreshFunc = func() tea.Msg { return nil }
	updated, cmd = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'r'}})
	next = updated.(ListModel[string])
	if !next.Refreshing || cmd == nil {
		t.Fatal("expected refresh action to trigger")
	}

	model = newTestListModel()
	model.EnableMerge = true
	updated, cmd = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'m'}})
	next = updated.(ListModel[string])
	if !next.MergeAction || cmd == nil {
		t.Fatal("expected merge action to trigger")
	}
}

func TestListModel_ActionsDoNotTriggerInSearchMode(t *testing.T) {
	model := newTestListModel()
	model.EnableYank = true
	model.EnableRefresh = true
	model.EnableMerge = true
	model.RefreshFunc = func() tea.Msg { return nil }

	updated, _ := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'/'}})
	model = updated.(ListModel[string])

	updated, cmd := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'y'}})
	next := updated.(ListModel[string])
	if cmd != nil {
		t.Fatal("did not expect yank command in search mode")
	}
	if next.YankAction {
		t.Fatal("did not expect yank action in search mode")
	}

	updated, cmd = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'r'}})
	next = updated.(ListModel[string])
	if cmd != nil || next.Refreshing {
		t.Fatal("did not expect refresh action in search mode")
	}

	updated, cmd = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'m'}})
	next = updated.(ListModel[string])
	if cmd != nil || next.MergeAction {
		t.Fatal("did not expect merge action in search mode")
	}
}
