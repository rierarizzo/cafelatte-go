package core

import "errors"

// User errors
var (
	ErrUnauthorizedUser = errors.New("unauthorized user")
)

// Generic errors
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrUnexpected     = errors.New("unexpected error")
	ErrBadRequest     = errors.New("bad request")
)

// Token errors
var (
	ErrSignAlgorithmUnexpected = errors.New("sign algorithm unexpected")
	ErrInvalidToken            = errors.New("invalid token")
	ErrParsingClaims           = errors.New("parsing claims error")
)
