package commands

import (
	"errors"
)

type ErrorCmd struct {
	message string
}

func NewErrorCmd(message string) ErrorCmd {
	return ErrorCmd{
		message: message,
	}
}

func (e ErrorCmd) Execute() error {
	return errors.New(e.message)
}
