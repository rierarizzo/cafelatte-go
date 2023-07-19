package errors

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrUnexpected     = errors.New("unexpected error")
	ErrBadRequest     = errors.New("bad request")
)
