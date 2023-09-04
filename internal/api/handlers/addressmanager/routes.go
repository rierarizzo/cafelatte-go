package addressmanager

import (
	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/usecases/addressmanager"
)

func Routes(g *echo.Group) func(m addressmanager.Manager) {
	addressGroup := g.Group("/address")

	return func(m addressmanager.Manager) {
		addressGroup.GET("/find/:userId", findAddressByUserId(m))
		addressGroup.POST("/register/:userId", registerAddressByUserId(m))
	}
}
