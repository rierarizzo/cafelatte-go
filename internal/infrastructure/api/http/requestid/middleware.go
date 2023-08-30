package requestid

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
)

func CustomMiddleware() echo.MiddlewareFunc {
	return middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return generateRequestId()
		},
	})
}

func generateRequestId() string {
	id := uuid.New()
	request.SetRequestId(id.String())

	return id.String()
}
