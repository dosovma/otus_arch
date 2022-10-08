package commands

import "errors"

var (
	ErrInvalidOperation = errors.New("invalid operation")
	ErrMacroCmd         = errors.New("command error")
)
