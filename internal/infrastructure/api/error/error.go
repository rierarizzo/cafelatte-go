package error

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/core/constants"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"net/http"
	"time"
)

type Response struct {
	HTTPStatus int       `json:"status"`
	Error      string    `json:"error"`
	Message    string    `json:"message"`
	RequestID  string    `json:"requestID"`
	TimeStamp  time.Time `json:"timestamp"`
}

func Error(c *gin.Context, err error) bool {
	if err != nil {
		_ = c.Error(err)

		errType, customMsg := utils.SeparateError(err)
		status := HTTPStatus(errType)

		requestID, _ := c.Get(constants.RequestIDKey)

		errorResponse := Response{
			HTTPStatus: status,
			Error:      errType.Error(),
			Message:    customMsg,
			RequestID:  requestID.(string),
			TimeStamp:  time.Now(),
		}

		c.AbortWithStatusJSON(status, errorResponse)
		return true
	}
	return false
}

func HTTPStatus(err error) int {
	errorStatusMap := map[error]int{
		errors.ErrInvalidUserFormat:       http.StatusBadRequest,
		errors.ErrInvalidAddressFormat:    http.StatusBadRequest,
		errors.ErrInvalidCardFormat:       http.StatusBadRequest,
		errors.ErrExpiredCard:             http.StatusBadRequest,
		errors.ErrTokenNotPresent:         http.StatusBadRequest,
		errors.ErrUnauthorizedUser:        http.StatusUnauthorized,
		errors.ErrInvalidToken:            http.StatusUnauthorized,
		errors.ErrSignAlgorithmUnexpected: http.StatusUnauthorized,
		errors.ErrRecordNotFound:          http.StatusNotFound,
	}

	if status, ok := errorStatusMap[err]; ok {
		return status
	}

	return http.StatusInternalServerError
}
