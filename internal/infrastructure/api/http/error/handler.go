package error

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
)

type Response struct {
	Status    int       `json:"status"`
	ErrorType string    `json:"errorType"`
	ErrorMsg  string    `json:"errorMsg"`
	IssuedAt  time.Time `json:"issuedAt"`
	RequestID any       `json:"requestID"`
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var appErr *domain.AppError
	ok := errors.As(err, &appErr)

	if ok {
		statusCode := getStatusCode(appErr.Type)
		errorResponse := Response{
			Status:    statusCode,
			ErrorType: appErr.Type,
			ErrorMsg:  appErr.Err.Error(),
			IssuedAt:  time.Now(),
			RequestID: request.ID(),
		}

		_ = c.JSON(statusCode, errorResponse)
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
