package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errs := c.Errors

		if len(errs) > 0 {
			var err *domain.AppError
			ok := errors.As(errs[0].Err, &err)
			if ok {
				if err.Type == domain.NotFoundError {
					c.JSON(http.StatusNotFound, err.Error())
					return
				} else if err.Type == domain.NotAuthorizedError || err.Type == domain.NotAuthenticatedError {
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
