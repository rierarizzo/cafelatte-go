package cardmanager

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/cardmanager"
)

func ConfigureRouting(g *echo.Group) func(m cardmanager.Manager) {
	cardsGroup := g.Group("/card")

	return func(m cardmanager.Manager) {
		cardsGroup.GET("/card/find/:userId", getCardsByUserId(m))
		cardsGroup.POST("/card/register/:userId", addNewCard(m))
	}
}

func getCardsByUserId(m cardmanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		cards, appErr := m.GetCardsByUserId(userId)
		if appErr != nil {
			return appErr
		}

		response := fromCardsToResponse(cards)
		return c.JSON(http.StatusOK, response)
	}
}

func addNewCard(m cardmanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		var request CardCreate
		err = c.Bind(&request)
		if err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		card, appErr := m.AddUserCard(userId, fromRequestToCard(request))
		if appErr != nil {
			return appErr
		}

		response := fromCardToResponse(card)
		return c.JSON(http.StatusCreated, response)
	}
}
