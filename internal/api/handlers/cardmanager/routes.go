package cardmanager

import (
	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/usecases/cardmanager"
)

func Routes(g *echo.Group) func(m cardmanager.Manager) {
	cardsGroup := g.Group("/card")

	return func(m cardmanager.Manager) {
		cardsGroup.GET("/card/find/:userId", getCardsByUserId(m))
		cardsGroup.POST("/card/register/:userId", addNewCard(m))
	}
}
