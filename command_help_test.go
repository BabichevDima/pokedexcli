package main

import (
	"testing"
)

func TestCommandHelp(t *testing.T) {
	err := commandHelp()

	if err != nil {
		t.Errorf("commandHelp() returned unexpected error: %v", err)
	}
}
