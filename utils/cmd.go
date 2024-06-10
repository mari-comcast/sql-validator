package utils

import (
	"errors"

	"github.com/comcast/sql-validator/schema"
)

func InterpretCommand(command []string) (schema.Command, error) {
	var cmd schema.Command

	for i, currCommand := range command {
		if currCommand == "--sql-file" || currCommand == "--text-file" {
			if i+1 >= len(command) {
				return cmd, errors.New("no file provided after " + currCommand)
			} else if command[i+1] == "" {
				return cmd, errors.New("empty file given after " + currCommand)
			}
		}

		switch currCommand {
		case "--sql-file":
			cmd.SqlFile = command[i+1]
		case "--text-file":
			cmd.TxtFile = command[i+1]
		}
	}

	if cmd.SqlFile == "" || cmd.TxtFile == "" {
		return cmd, errors.New("empty sql file or value file given")
	}

	return cmd, nil
}
