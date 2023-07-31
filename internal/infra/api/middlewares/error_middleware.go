package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/params"
	"net/http"
	"time"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors[0].Err
			var appErr *domain.AppError
			ok := errors.As(err, &appErr)
			if ok {
				if appErr.Type == domain.NotFoundError {
					writeError(c, http.StatusNotFound, err)
					return
				} else if appErr.Type == domain.NotAuthorizedError ||
					appErr.Type == domain.NotAuthenticatedError ||
					appErr.Type == domain.TokenValidationError {
					writeError(c, http.StatusUnauthorized, err)
					return
				} else {
					writeError(c, http.StatusInternalServerError, err)
					return
				}
			}

			writeError(c, http.StatusInternalServerError, err)
			return
		}
	}
}

type ErrorResponse struct {
	Status    int       `json:"status"`
	ErrorType string    `json:"errorType"`
	ErrorMsg  string    `json:"errorMsg"`
	IssuedAt  time.Time `json:"issuedAt"`
	RequestID any       `json:"requestID"`
}

func writeError(c *gin.Context, httpStatus int, err error) {
	var appErr *domain.AppError
	converted := errors.As(err, &appErr)

	if !converted {
		appErr = domain.NewAppError(err, domain.UnexpectedError)
	}

	response := ErrorResponse{
		Status:    httpStatus,
		ErrorType: appErr.Type,
		ErrorMsg:  appErr.Err.Error(),
		IssuedAt:  time.Now(),
		RequestID: params.RequestID(),
	}

	c.JSON(httpStatus, response)
}
