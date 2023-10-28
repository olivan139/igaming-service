package errs

import "errors"

var (
	ErrOutOfRange      = errors.New("reels index out of range")
	ErrLineNotFound    = errors.New("line not found")
	ErrUndefinedSymbol = errors.New("undefined symbol")
)
