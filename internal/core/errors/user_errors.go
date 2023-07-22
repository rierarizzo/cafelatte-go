package errors

import "errors"

var (
	ErrUnauthorizedUser     = errors.New("unauthorized user")
	ErrInvalidUserFormat    = errors.New("invalid user format")
	ErrExpiredCard          = errors.New("payment card is expired")
	ErrInvalidCardFormat    = errors.New("invalid payment card format")
	ErrInvalidAddressFormat = errors.New("invalid address format")
)
