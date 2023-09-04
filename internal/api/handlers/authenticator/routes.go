package authenticator

import (
	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/usecases/authenticator"
)

func Routes(group *echo.Group) func(a authenticator.Authenticator) {
	return func(a authenticator.Authenticator) {
		group.POST("/signup", signUp(a))
		group.POST("/signin", signIn(a))
	}
}
