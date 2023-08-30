package cardmanager

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/cardmanager"
)

type Handler struct {
	cardManager cardmanager.Manager
}

func Router(group *echo.Group) func(cardManagerHandler *Handler) {
	return func(h *Handler) {
		group.GET("/card/find/:userId", h.GetCardsByUserId)
		group.POST("/card/register/:userId", h.AddCard)
	}
}

func (h *Handler) GetCardsByUserId(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	cards, appErr := h.cardManager.GetCardsByUserId(userId)
	if appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusOK, fromCardsToResponse(cards))
}

func (h *Handler) AddCard(c echo.Context) error {
	var req CardCreate
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	if err = c.Bind(&req); err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	card, appErr := h.cardManager.AddUserCard(userId, fromRequestToCard(req))
	if appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusCreated, fromCardToResponse(card))
}

func New(cardManager cardmanager.Manager) *Handler {
	return &Handler{cardManager}
}
