package builtins

import (
	"bytes"
	"os"
	"testing"
)

func TestPrintWorkingDirectory(t *testing.T) {
	// Create a temporary directory for testing.
	tempDir, err := os.MkdirTemp("", "testpwd")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Change the working directory to the temporary directory.
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	defer func() {
		if chdirErr := os.Chdir(oldDir); chdirErr != nil {
			t.Errorf("Failed to restore working directory: %v", chdirErr)
		}
	}()

	// Attempt to change the working directory.
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change working directory: %v", err)
	}

	// Create a buffer to capture the output.
	var buf bytes.Buffer

	// Call the PrintWorkingDirectory function.
	err = PrintWorkingDirectory(&buf)
	if err != nil {
		t.Errorf("PrintWorkingDirectory failed with error: %v", err)
	}

	// Check if the output matches the temporary directory path.
	expectedOutput := tempDir + "\n"
	if buf.String() != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, buf.String())
	}
}
