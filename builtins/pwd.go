package builtins

import (
	"fmt"
	"io"
	"os"
)

// PrintWorkingDirectory prints the current working directory.
func PrintWorkingDirectory(w io.Writer) error {
	// Get the current working directory.
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(w, currentDir)
	return err
}
