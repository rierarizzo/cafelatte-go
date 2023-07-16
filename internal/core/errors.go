package core

import "errors"

// User errors
var (
	UnauthorizedUser = errors.New("unauthorized user")
)

// Generic errors
var (
	RecordNotFound = errors.New("record not found")
	Unexpected     = errors.New("unexpected error")
)
