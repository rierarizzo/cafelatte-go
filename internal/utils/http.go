package utils

import (
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HTTPError(coreError error, c *gin.Context) {
	logrus.Error(coreError)
	switch coreError {
	case errors.ErrBadRequest:
		c.AbortWithStatusJSON(http.StatusBadRequest, coreError.Error())
	case errors.ErrRecordNotFound:
		c.AbortWithStatusJSON(http.StatusNotFound, coreError.Error())
	case errors.ErrUnauthorizedUser:
		c.AbortWithStatusJSON(http.StatusUnauthorized, coreError.Error())
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, coreError.Error())
	}
}
