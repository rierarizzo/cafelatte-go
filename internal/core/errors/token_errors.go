package errors

import "errors"

var (
	ErrSignAlgorithmUnexpected = errors.New("sign algorithm unexpected")
	ErrInvalidToken            = errors.New("invalid token")
	ErrTokenNotPresent         = errors.New("token not present")
)
