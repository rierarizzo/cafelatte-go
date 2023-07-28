package errors

import (
	"errors"
	"fmt"
	"slices"
	"testing"
)

const (
	errMsg1 = "error msg 1"
	errMsg2 = "error msg 2"
)

func TestWrapError(t *testing.T) {
	validate := func(t *testing.T, result error, expected error) {
		if result.Error() != expected.Error() {
			t.Errorf(
				"Result was incorrect, got: %s, want: %s.",
				result,
				expected,
			)
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

func TestSplitError(t *testing.T) {
	wrappedErr := WrapError(ErrUnexpected, errMsg1)
	wrappedErr = WrapError(wrappedErr, errMsg2)

	var expectedErrMsgs []string
	expectedErrMsgs = append(expectedErrMsgs, errMsg1)
	expectedErrMsgs = append(expectedErrMsgs, errMsg2)

	resultCoreErr, resultErrMsgs := SplitError(wrappedErr)

	if !errors.Is(resultCoreErr, ErrUnexpected) {
		t.Errorf(
			"Result was incorrect, got: %v, want: %v.",
			resultCoreErr,
			ErrUnexpected,
		)
	}

	if slices.Compare(resultErrMsgs, expectedErrMsgs) != 0 {
		t.Errorf(
			"Result was incorrect, got: %v, want: %v.",
			resultErrMsgs,
			expectedErrMsgs,
		)
	}
}
