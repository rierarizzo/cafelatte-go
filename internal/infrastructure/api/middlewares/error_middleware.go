package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
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
					writeError(c, http.StatusNotFound, appErr)
					return
				} else if appErr.Type == domain.NotAuthorizedError ||
					appErr.Type == domain.NotAuthenticatedError ||
					appErr.Type == domain.TokenValidationError {
					writeError(c, http.StatusUnauthorized, appErr)
					return
				} else if appErr.Type == domain.BadRequestError {
					writeError(c, http.StatusBadRequest, appErr)
				} else {
					writeError(c, http.StatusInternalServerError, appErr)
					return
				}
			}

			writeError(c, http.StatusInternalServerError,
				domain.NewAppError(err, domain.UnexpectedError))
			return
		}
	}
}

func writeError(c *gin.Context, httpStatus int, appErr *domain.AppError) {
	response := dto.ErrorResponse{
		Status:    httpStatus,
		ErrorType: appErr.Type,
		ErrorMsg:  appErr.Err.Error(),
		IssuedAt:  time.Now(),
		RequestID: params.RequestID(),
	}

	c.JSON(httpStatus, response)
}
