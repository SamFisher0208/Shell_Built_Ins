package builtins

import (
	"os"
	"testing"
)

func TestMakeDirectory(t *testing.T) {
	// Define a temporary directory name for testing.
	testDirName := "testdir"

	// Clean up any remnants from previous tests.
	defer func() {
		_ = os.Remove(testDirName)
	}()

	// Test case 1: Successfully create a directory.
	err := MakeDirectory(testDirName)
	if err != nil {
		t.Errorf("MakeDirectory failed with error: %v", err)
	}

	// Check if the directory exists.
	if _, err := os.Stat(testDirName); os.IsNotExist(err) {
		t.Errorf("MakeDirectory should have created the directory, but it doesn't exist")
	}

	// Test case 2: Attempt to create the same directory again (should return an error).
	err = MakeDirectory(testDirName)
	if err == nil {
		t.Errorf("MakeDirectory should have returned an error for existing directory, but it didn't")
	}
}
