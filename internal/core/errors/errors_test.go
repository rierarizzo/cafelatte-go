package errors

import (
	"fmt"
	"testing"
)

func TestWrapError(t *testing.T) {
	validate := func(t *testing.T, result error, expected error) {
		if result.Error() != expected.Error() {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, expected)
		}
	}

	format := "%w: %s"

	errMsg := "an unexpected error has ocurred"
	expected := fmt.Errorf(format, ErrUnexpected, errMsg)
	result := WrapError(ErrUnexpected, errMsg)

	validate(t, result, expected)

	errMsg = "another message"
	expected = fmt.Errorf(format, result, errMsg)
	result = WrapError(result, errMsg)

	validate(t, result, expected)
}

func TestCompareErrors(t *testing.T) {
	validate := func(t *testing.T, wrappedErr error) {
		if !CompareErrors(wrappedErr, ErrUnexpected) {
			t.Errorf("Result was incorrect, errors are not equals.")
		}
	}

	wrappedErr := WrapError(ErrUnexpected, "error msg 1")
	validate(t, wrappedErr)

	wrappedErr = WrapError(wrappedErr, "error msg 2")
	validate(t, wrappedErr)
}
