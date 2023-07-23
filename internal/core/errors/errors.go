package errors

import (
	"errors"
	"fmt"
	"strings"
)

const errorMsgFormat = "%w: %s"

var (
	ErrUnauthorizedUser     = errors.New("UnauthorizedUserError")
	ErrInvalidUserFormat    = errors.New("InvalidUserFormatError")
	ErrExpiredCard          = errors.New("ExpiredCardError")
	ErrInvalidCardFormat    = errors.New("InvalidCardFormatError")
	ErrInvalidAddressFormat = errors.New("InvalidAddressFormatError")
)

var (
	ErrSignAlgorithmUnexpected = errors.New("SignAlgorithmUnexpectedError")
	ErrInvalidToken            = errors.New("InvalidTokenError")
	ErrTokenNotPresent         = errors.New("TokenNotPresentError")
)

var (
	ErrRecordNotFound = errors.New("RecordNotFoundError")
	ErrUnexpected     = errors.New("UnexpectedError")
)

func WrapError(errType error, errReturned string) error {
	errReturned = strings.ReplaceAll(errReturned, ":", ",")
	return fmt.Errorf(errorMsgFormat, errType, errReturned)
}
