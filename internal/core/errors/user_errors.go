package errors

import "errors"

var (
	ErrUnauthorizedUser = errors.New("unauthorized user")
	ErrInvalidUserData  = errors.New("invalid user data")
)
