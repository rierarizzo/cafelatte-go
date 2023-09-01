package error

import (
	"errors"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type Body struct {
	ErrorType string    `json:"errorType"`
	ErrorMsg  string    `json:"errorMsg"`
	IssuedAt  time.Time `json:"issuedAt"`
}

func CustomHttpErrorHandler(err error, c echo.Context) {
	var appErr *domain.AppError
	ok := errors.As(err, &appErr)

	if ok {
		statusCode := getStatusCode(appErr.Type)
		errorBody := Body{
			ErrorType: appErr.Type,
			ErrorMsg:  appErr.Err.Error(),
			IssuedAt:  time.Now(),
		}

		_ = httpUtil.RespondWithError(c, statusCode, errorBody)
	}
}

func getStatusCode(appErrorType string) int {
	errorStatusCodeMaps := map[string]int{
		domain.NotFoundError:         http.StatusNotFound,
		domain.NotAuthorizedError:    http.StatusUnauthorized,
		domain.NotAuthenticatedError: http.StatusUnauthorized,
		domain.TokenValidationError:  http.StatusUnauthorized,
		domain.BadRequestError:       http.StatusBadRequest,
	}

	for key, value := range errorStatusCodeMaps {
		if appErrorType == key {
			return value
		}
	}

	return http.StatusInternalServerError
}
