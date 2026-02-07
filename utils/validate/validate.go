// Package validate provides validation functions for user input.
package validate

import (
	"fmt"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// Username validates that a username is non-empty and contains no whitespace.
func Username(username string) error {
	if username == "" {
		return fmt.Errorf("username cannot be empty")
	}
	if strings.TrimSpace(username) != username {
		return fmt.Errorf("username cannot have leading or trailing whitespace")
	}
	if strings.ContainsAny(username, " \t\n\r") {
		return fmt.Errorf("username cannot contain whitespace")
	}
	return nil
}

// Token validates token format. Empty token is allowed (optional field).
func Token(token string) error {
	if token == "" {
		return nil // Token is optional
	}
	if strings.TrimSpace(token) != token {
		return fmt.Errorf("token cannot have leading or trailing whitespace")
	}
	if strings.ContainsAny(token, " \t\n\r") {
		return fmt.Errorf("token cannot contain whitespace")
	}
	return nil
}

// Email validates email format. Empty email is allowed (optional field).
func Email(email string) error {
	if email == "" {
		return nil
	}
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}
	return nil
}
