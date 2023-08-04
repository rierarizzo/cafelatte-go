package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/constants"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
	"github.com/rierarizzo/cafelatte/internal/params"
	"github.com/sirupsen/logrus"
	"time"
)

func AbortWithError(c *gin.Context, appErr *domain.AppError) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	log.Error(c.Error(appErr))
	c.Abort()
}

func RespondWithJSON(c *gin.Context, statusCode int, body interface{}) {
	response := dto.OKResponse{
		Status:    statusCode,
		Body:      body,
		IssuedAt:  time.Now(),
		RequestID: params.RequestID(),
	}

	c.JSON(statusCode, response)
}
