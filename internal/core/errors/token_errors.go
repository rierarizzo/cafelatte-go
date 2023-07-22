package errors

import "errors"

var (
	ErrSignAlgorithmUnexpected = errors.New("SignAlgorithmUnexpectedError")
	ErrInvalidToken            = errors.New("InvalidTokenError")
	ErrTokenNotPresent         = errors.New("TokenNotPresentError")
)
