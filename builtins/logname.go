package builtins

import (
	"fmt"
	"io"
	"os/user"
)

// Logname prints the current user's login name.
func Logname(w io.Writer, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("%w: expected no arguments", ErrInvalidArgCount)
	}

	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	loginName := currentUser.Username
	_, err = fmt.Fprintln(w, loginName)
	return err
}
