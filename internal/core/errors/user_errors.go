package errors

import "errors"

var (
	ErrUnauthorizedUser     = errors.New("UnauthorizedUserError")
	ErrInvalidUserFormat    = errors.New("InvalidUserFormatError")
	ErrExpiredCard          = errors.New("ExpiredCardError")
	ErrInvalidCardFormat    = errors.New("InvalidCardFormatError")
	ErrInvalidAddressFormat = errors.New("InvalidAddressFormatError")
)
