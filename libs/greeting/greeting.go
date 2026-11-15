// Package greeting builds the greetings the greeter service returns.
package greeting

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// MaxNameLength is the longest name, in characters, that Greet accepts.
const MaxNameLength = 64

// ErrEmptyName means the name holds no visible character.
var ErrEmptyName = errors.New("greeting: name is empty")

// ErrNameTooLong means the name is longer than MaxNameLength characters.
var ErrNameTooLong = errors.New("greeting: name is too long")

// Greet returns "Hello, <name>!" for a trimmed, non-empty name.
func Greet(name string) (string, error) {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return "", ErrEmptyName
	}
	if utf8.RuneCountInString(trimmed) > MaxNameLength {
		return "", ErrNameTooLong
	}
	return "Hello, " + trimmed + "!", nil
}
