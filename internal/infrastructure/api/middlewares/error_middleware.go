package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	core "github.com/rierarizzo/cafelatte/internal/core/errors"
	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errs := c.Errors

		if len(errs) > 0 {
			var err *core.AppError
			ok := errors.As(errs[0].Err, &err)
			if ok {
				if err.Type == core.NotFoundError {
					c.JSON(http.StatusNotFound, err.Error())
					return
				} else if err.Type == core.NotAuthorizedError || err.Type == core.NotAuthenticatedError {
					c.JSON(http.StatusUnauthorized, err.Error())
					return
				} else {
					c.JSON(http.StatusInternalServerError, err.Error())
					return
				}
			}

			c.JSON(http.StatusInternalServerError, "internal server error")
			return
		}
	}
}
