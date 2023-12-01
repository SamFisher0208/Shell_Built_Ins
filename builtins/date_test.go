package builtins

import (
	"bytes"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	// Create a buffer to capture the output.
	var buf bytes.Buffer

	// Set the expected format for the date and time.
	expectedFormat := "Mon Jan 2 03:04 PM MST 2006"

	// Define a function that returns a fixed time for testing.
	getFixedTime := func() time.Time {
		return time.Date(2023, time.February, 15, 14, 30, 0, 0, time.UTC)
	}

	// Call the Date function with the custom time function.
	err := Date(&buf, getFixedTime)
	if err != nil {
		t.Errorf("Date failed with error: %v", err)
	}

	// Check if the output matches the expected format.
	if buf.String() != getFixedTime().Format(expectedFormat)+"\n" {
		t.Errorf("Expected output: %s, but got: %s", getFixedTime().Format(expectedFormat), buf.String())
	}
}
