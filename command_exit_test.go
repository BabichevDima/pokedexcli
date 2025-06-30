package main

import (
	"bytes"
	"os"
	"testing"
)

var gotExit bool
var exitCode int

func fakeExit(code int) {
	gotExit = true
	exitCode = code
}

func TestCommandExit(t *testing.T) {
	oldExit := osExit
	osExit = fakeExit
	defer func() { osExit = oldExit }()

	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = oldOut }()

	err := commandExit()

	w.Close()
	os.Stdout = oldOut

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expectedMsg := "Closing the Pokedex... Goodbye!\n"
	if output != expectedMsg {
		t.Errorf("Expected output %q, got %q", expectedMsg, output)
	}

	if !gotExit {
		t.Error("Expected os.Exit to be called")
	}

	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}

	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}