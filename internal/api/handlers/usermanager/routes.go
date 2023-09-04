package usermanager

import (
	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/usecases/usermanager"
)

func Routes(g *echo.Group) func(m usermanager.Manager) {
	return func(m usermanager.Manager) {
		g.GET("/find", getAllUsers(m))
		g.GET("/find/:userId", findUserById(m))
		g.PUT("/update/:userId", updateUserById(m))
		g.DELETE("/delete/:userId", deleteUserById(m))
	}
}
