package builtins

import (
	"fmt"
	"io"
	"os"
)

// PrintWorkingDirectory prints the current working directory.
func PrintWorkingDirectory(w io.Writer, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("%w: expected no arguments", ErrInvalidArgCount)
	}

	// Get the current working directory.
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(w, currentDir)
	return err
}
