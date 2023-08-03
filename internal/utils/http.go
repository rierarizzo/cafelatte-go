package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/constants"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/params"
	"github.com/sirupsen/logrus"
)

func AbortWithError(c *gin.Context, appErr *domain.AppError) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	log.Error(c.Error(appErr))
	c.Abort()
}
