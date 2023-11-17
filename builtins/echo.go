package builtins

import (
	"fmt"
	"io"
	"strings"
)

// Echo echoes its arguments to the specified writer on a single line.
func Echo(w io.Writer, args ...string) error {
	echoStr := strings.Join(args, " ")
	_, err := fmt.Fprintln(w, echoStr)
	return err
}
