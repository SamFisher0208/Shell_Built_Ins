package builtins

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	// Create a buffer to capture the output.
	var buf bytes.Buffer

	// Test case 1: Check if "Hello, World!" is echoed correctly.
	err := Echo(&buf, "Hello,", "World!")
	if err != nil {
		t.Errorf("Echo failed with error: %v", err)
	}
	expectedOutput := "Hello, World!\n"
	if buf.String() != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, buf.String())
	}

}
