package builtins

import (
	"fmt"
	"os"
)

func MakeDirectory(args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("%w: expected one argument (directory name)", ErrInvalidArgCount)
	}

	dirName := args[0]

	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
