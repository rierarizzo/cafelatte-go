package http

import (
	"github.com/gin-gonic/gin"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
	"time"
)

func AbortWithError(c *gin.Context, appErr *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	log.Error(c.Error(appErr))
	c.Abort()
}

type OKResponse struct {
	Status    int         `json:"status"`
	Body      interface{} `json:"body"`
	IssuedAt  time.Time   `json:"issuedAt"`
	RequestID any         `json:"requestID"`
}

func RespondWithJSON(c *gin.Context, statusCode int, body interface{}) {
	response := OKResponse{
		Status:    statusCode,
		Body:      body,
		IssuedAt:  time.Now(),
		RequestID: request.ID(),
	}

	c.JSON(statusCode, response)
}
