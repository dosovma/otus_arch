package commands

import "errors"

var (
	ErrInvalidOperation  = errors.New("invalid operation")
	ErrConnectionTimeout = errors.New("connection timeout")
)
