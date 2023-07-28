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
	return fmt.Errorf(errorMsgFormat, errType, errReturned)
}

func GetOriginalErrorType(err error) error {
	for {
		if wrapped, ok := err.(interface{ Unwrap() error }); ok {
			err = wrapped.Unwrap()
		} else {
			break
		}
	}
	return err
}

func SplitError(err error) (error, []string) {
	errorMsgs := strings.Split(err.Error(), ":")[1:]
	for k, v := range errorMsgs {
		errorMsgs[k] = strings.TrimSpace(v)
	}

	return GetOriginalErrorType(err), errorMsgs
}
