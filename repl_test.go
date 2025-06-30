package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	var cases = []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		
		if len(actual) != len(c.expected) {
			t.Errorf("Length mismatch for input '%s': expected %d, got %d", c.input, len(c.expected), len(actual))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Word mismatch for input '%s' at position %d: expected '%s', got '%s'", c.input, i, c.expected[i], actual[i])
			}
		}
	}
}
 
func TestGetCommands(t *testing.T) {
	type cliCommand struct {
		name        string
		description string
		callback    func() error
	}

	var expected = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	actual := getCommands()

	if len(actual) != len(expected) {
		t.Errorf("Length mismatch: expected %d, got %d", len(expected), len(actual))
	}
}