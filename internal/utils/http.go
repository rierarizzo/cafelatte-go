package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/constants"
	"github.com/rierarizzo/cafelatte/internal/params"
	"github.com/sirupsen/logrus"
)

func AbortWithError(c *gin.Context, err error) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	log.Error(c.Error(err))
	c.Abort()
}
