package builtins

import (
	"fmt"
	"io"
	"time"
)

// Date prints the current date and time, formatted to 12 hour.
func Date(w io.Writer, funcTime func() time.Time, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("%w: expected no arguments", ErrInvalidArgCount)
	}

	currentTime := funcTime()
	dateStr := currentTime.Format("Mon Jan 2 03:04 PM MST 2006")
	_, err := fmt.Fprintln(w, dateStr)
	return err
}
