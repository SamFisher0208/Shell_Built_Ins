package builtins_test

import (
	"bytes"
	"testing"

	"github.com/SamFisher0208/Shell_Built_Ins/builtins"
)

func TestLogname(t *testing.T) {
	// Create a buffer to capture the output.
	var buf bytes.Buffer

	// Call the Logname function.
	err := builtins.Logname(&buf)
	if err != nil {
		t.Errorf("Logname failed with error: %v", err)
	}

	// Check the captured login name.
	expectedName := "IRONHIDE\\Test User"
	actualName := buf.String()
	if actualName == expectedName+"\n" {
		t.Errorf("Expected login name to be different from: %s, but it matches: %s", expectedName, actualName)
	}
}
