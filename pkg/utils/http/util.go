package http

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Success   bool        `json:"success"`
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	RequestId string      `json:"requestId"`
}

func RespondWithJSON(c echo.Context, code int, data interface{}) error {
	response := Response{
		Success:   true,
		Code:      code,
		Data:      data,
		RequestId: c.Request().Header.Get("X-Request-ID"),
	}

	return c.JSON(code, response)
}

func RespondWithError(c echo.Context, code int, data interface{}) error {
	response := Response{
		Success:   false,
		Code:      code,
		Data:      data,
		RequestId: c.Request().Header.Get("X-Request-ID"),
	}

	return c.JSON(code, response)
}
