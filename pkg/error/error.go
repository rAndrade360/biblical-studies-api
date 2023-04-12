package error

import "errors"

var (
	NOTFOUND     = errors.New("not found")
	INVALIDINPUT = errors.New("invalid input")
)
