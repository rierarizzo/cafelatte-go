package productmanager

import (
	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/usecases/productmanager"
)

func Routes(g *echo.Group) func(m productmanager.Manager) {
	return func(m productmanager.Manager) {
		g.GET("/find", getAllProducts(m))
		g.GET("/find/categories", getProductCategories(m))
	}
}
