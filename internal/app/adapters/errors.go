package adapters

import "errors"

var (
	ErrNotImplemented  = errors.New("operation is not implemented")
	ErrInvalidProperty = errors.New("received invalid property")
)
