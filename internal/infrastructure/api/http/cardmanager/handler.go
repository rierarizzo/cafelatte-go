package cardmanager

import (
	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/cardmanager"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/authenticator"
	"net/http"
	"strconv"
)

type Handler struct {
	cardManager cardmanager.Manager
}

func Router(group *echo.Group) func(cardManagerHandler *Handler) {
	return func(handler *Handler) {
		cardsGroup := group.Group("/cards")

		cardsGroup.Use(authenticator.Middleware)

		cardsGroup.GET("/find/:userId", handler.GetCardsByUserId)
		cardsGroup.POST("/register/:userId", handler.AddUserCards)
	}
}

func (handler *Handler) GetCardsByUserId(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	cards, appErr := handler.cardManager.GetCardsByUserId(userId)
	if appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusOK, fromCardsToResponse(cards))
}

func (handler *Handler) AddUserCards(c echo.Context) error {
	var req RegisterCardRequest
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	if err := c.Bind(&req); err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	card, appErr := handler.cardManager.AddUserCard(userId, fromRequestToCard(req))
	if appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusCreated, fromCardToResponse(card))
}

func New(paymentCardService cardmanager.Manager) *Handler {
	return &Handler{paymentCardService}
}
