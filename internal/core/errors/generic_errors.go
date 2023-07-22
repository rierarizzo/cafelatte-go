package errors

import (
	"errors"
	"fmt"
	"strings"
)

const errorMsgFormat = "%w: %s"

var (
	ErrRecordNotFound = errors.New("RecordNotFoundError")
	ErrUnexpected     = errors.New("UnexpectedError")
)

func WrapError(errType error, errReturned string) error {
	errReturned = strings.ReplaceAll(errReturned, ":", ",")
	return fmt.Errorf(errorMsgFormat, errType, errReturned)
}
