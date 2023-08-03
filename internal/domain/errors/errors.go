package errors

import (
	"errors"
)

const (
	NotFoundError   = "NotFoundError"
	notFoundMessage = "record not found"

	ValidationError        = "ValidationError"
	validationErrorMessage = "validation errors"

	ResourceAlreadyExistsError        = "ResourceAlreadyExistsError"
	resourceAlreadyExistsErrorMessage = "resource already exists"

	RepositoryError        = "RepositoryError"
	repositoryErrorMessage = "errors in repository operation"

	NotAuthenticatedError        = "NotAuthenticatedError"
	notAuthenticatedErrorMessage = "not authenticated"

	TokenGenerationError        = "TokenGenerationError"
	tokenGenerationErrorMessage = "errors in token generation"

	TokenValidationError        = "TokenValidationError"
	tokenValidationErrorMessage = "errors in token validation"

	HashGenerationError        = "HashGenerationError"
	hashGenerationErrorMessage = "errors in hash generation"

	NotAuthorizedError        = "NotAuthorizedError"
	notAuthorizedErrorMessage = "not authorized"

	DuplicatedResourceError        = "DuplicatedResourceError"
	duplicatedResourceErrorMessage = "resource is duplicated"

	BadRequestError        = "BadRequestError"
	badRequestErrorMessage = "bad request"

	UnexpectedError        = "UnexpectedError"
	unexpectedErrorMessage = "something went wrong"
)

type AppError struct {
	Err  error
	Type string
}

// NewAppError returns an AppError with its type and message.
func NewAppError(err error, errType string) *AppError {
	return &AppError{
		Err:  err,
		Type: errType,
	}
}

// NewAppErrorWithType returns an AppError with its respective type.
func NewAppErrorWithType(errType string) *AppError {
	var err error

	switch errType {
	case NotFoundError:
		err = errors.New(notFoundMessage)
	case ValidationError:
		err = errors.New(validationErrorMessage)
	case ResourceAlreadyExistsError:
		err = errors.New(resourceAlreadyExistsErrorMessage)
	case RepositoryError:
		err = errors.New(repositoryErrorMessage)
	case NotAuthenticatedError:
		err = errors.New(notAuthenticatedErrorMessage)
	case TokenGenerationError:
		err = errors.New(tokenGenerationErrorMessage)
	case TokenValidationError:
		err = errors.New(tokenValidationErrorMessage)
	case HashGenerationError:
		err = errors.New(hashGenerationErrorMessage)
	case NotAuthorizedError:
		err = errors.New(notAuthorizedErrorMessage)
	case DuplicatedResourceError:
		err = errors.New(duplicatedResourceErrorMessage)
	case BadRequestError:
		err = errors.New(badRequestErrorMessage)
	case UnexpectedError:
		err = errors.New(unexpectedErrorMessage)
	}

	return &AppError{
		Err:  err,
		Type: errType,
	}
}

// Error converts the app errors to a human-readable text.
func (appErr *AppError) Error() string {
	return appErr.Err.Error()
}
