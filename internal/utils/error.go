package utils

import "strings"

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

func SeparateError(err error) (error, string) {
	splitedErr := strings.Split(err.Error(), ":")

	errCustomMsg := ""
	if len(splitedErr) > 1 {
		errCustomMsg = strings.TrimSpace(splitedErr[1])
	}

	return GetOriginalErrorType(err), errCustomMsg
}

func CompareErrors(err error, coreError error) bool {
	return GetOriginalErrorType(err) == coreError
}
