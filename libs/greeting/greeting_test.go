package greeting

import (
	"errors"
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{name: "plain name", input: "Ada", want: "Hello, Ada!"},
		{name: "trims spaces", input: "  Grace  ", want: "Hello, Grace!"},
		{name: "keeps unicode", input: "Zoé", want: "Hello, Zoé!"},
		{name: "empty", input: "", wantErr: ErrEmptyName},
		{name: "blank", input: " \t ", wantErr: ErrEmptyName},
		{name: "longest accepted", input: strings.Repeat("é", MaxNameLength), want: "Hello, " + strings.Repeat("é", MaxNameLength) + "!"},
		{name: "too long", input: strings.Repeat("a", MaxNameLength+1), wantErr: ErrNameTooLong},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Greet(tc.input)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Greet(%q) error = %v, want %v", tc.input, err, tc.wantErr)
			}
			if got != tc.want {
				t.Fatalf("Greet(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestFarewell(t *testing.T) {
	got, err := Farewell("  Ada ")
	if err != nil || got != "Goodbye, Ada!" {
		t.Fatalf("Farewell(%q) = %q, %v, want %q, nil", "  Ada ", got, err, "Goodbye, Ada!")
	}
	if _, err := Farewell(""); !errors.Is(err, ErrEmptyName) {
		t.Fatalf("Farewell(\"\") error = %v, want %v", err, ErrEmptyName)
	}
}
