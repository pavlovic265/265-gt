package validate

import (
	"testing"
)

func TestUsername_Valid(t *testing.T) {
	validUsernames := []string{"user", "user123", "user-name", "user_name", "user.name"}
	for _, username := range validUsernames {
		if err := Username(username); err != nil {
			t.Errorf("Username(%q) should be valid, got error: %v", username, err)
		}
	}
}

func TestUsername_Empty(t *testing.T) {
	err := Username("")
	if err == nil {
		t.Error("Username(\"\") should return error")
	}
}

func TestUsername_WithWhitespace(t *testing.T) {
	invalidUsernames := []string{" user", "user ", " user ", "user name", "user\tname", "user\nname"}
	for _, username := range invalidUsernames {
		if err := Username(username); err == nil {
			t.Errorf("Username(%q) should return error for whitespace", username)
		}
	}
}

func TestToken_Valid(t *testing.T) {
	validTokens := []string{"ghp_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "glpat-xxxxxxxxxxxxxxxxxxxx", "abc123"}
	for _, token := range validTokens {
		if err := Token(token); err != nil {
			t.Errorf("Token(%q) should be valid, got error: %v", token, err)
		}
	}
}

func TestToken_Empty(t *testing.T) {
	err := Token("")
	if err != nil {
		t.Error("Token(\"\") should be valid (optional field)")
	}
}

func TestToken_WithWhitespace(t *testing.T) {
	invalidTokens := []string{" token", "token ", " token ", "tok en"}
	for _, token := range invalidTokens {
		if err := Token(token); err == nil {
			t.Errorf("Token(%q) should return error for whitespace", token)
		}
	}
}

func TestEmail_Valid(t *testing.T) {
	validEmails := []string{"user@example.com", "user.name@example.co.uk", "user+tag@example.org"}
	for _, email := range validEmails {
		if err := Email(email); err != nil {
			t.Errorf("Email(%q) should be valid, got error: %v", email, err)
		}
	}
}

func TestEmail_Empty(t *testing.T) {
	err := Email("")
	if err != nil {
		t.Error("Email(\"\") should be valid (optional field)")
	}
}

func TestEmail_Invalid(t *testing.T) {
	invalidEmails := []string{"user", "user@", "@example.com", "user@example", "user example.com"}
	for _, email := range invalidEmails {
		if err := Email(email); err == nil {
			t.Errorf("Email(%q) should return error for invalid format", email)
		}
	}
}
