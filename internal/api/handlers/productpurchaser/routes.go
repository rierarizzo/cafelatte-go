package productpurchaser

import (
	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/usecases/productpurchaser"
)

func Routes(g *echo.Group) func(p productpurchaser.Purchaser) {
	return func(p productpurchaser.Purchaser) {
		g.POST("/", purchase(p))
	}
}
