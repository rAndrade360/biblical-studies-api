package error

import "errors"

var (
	NOTFOUND                  = errors.New("not found")
	INVALIDINPUT              = errors.New("invalid input")
	ITERNAL_SERVER_ERROR_HTTP = errors.New(`{"message": "internal server error"}`)
	BAD_REQUEST_HTTP          = errors.New(`{"message": "bad request"}`)
)
