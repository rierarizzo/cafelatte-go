package error

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"net/http"
	"time"
)

func Middleware() gin.HandlerFunc {
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

type Response struct {
	Status    int       `json:"status"`
	ErrorType string    `json:"errorType"`
	ErrorMsg  string    `json:"errorMsg"`
	IssuedAt  time.Time `json:"issuedAt"`
	RequestID any       `json:"requestID"`
}

func writeError(c *gin.Context, httpStatus int, appErr *domain.AppError) {
	response := Response{
		Status:    httpStatus,
		ErrorType: appErr.Type,
		ErrorMsg:  appErr.Err.Error(),
		IssuedAt:  time.Now(),
		RequestID: request.ID(),
	}

	c.JSON(httpStatus, response)
}
