package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/core"
	"github.com/sirupsen/logrus"
)

func HTTPError(coreError error, c *gin.Context) {
	logrus.Error(coreError)
	switch coreError {
	case core.BadRequest:
		c.AbortWithStatusJSON(http.StatusBadRequest, coreError.Error())
	case core.RecordNotFound:
		c.AbortWithStatusJSON(http.StatusNotFound, coreError.Error())
	case core.UnauthorizedUser:
		c.AbortWithStatusJSON(http.StatusUnauthorized, coreError.Error())
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, coreError.Error())
	}
}
