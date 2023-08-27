package saverequestid

import (
	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/pkg/constants/http"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
)

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		request.SetRequestID(c.Response().Header().Get(http.RequestIDHeader))
		return next(c)
	}
}
