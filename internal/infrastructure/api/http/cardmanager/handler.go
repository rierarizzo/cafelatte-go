package cardmanager

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/cardmanager"
	http2 "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"
	"strconv"
)

type Handler struct {
	paymentCardService cardmanager.Manager
}

func (h *Handler) GetCardsByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		http2.AbortWithError(c, appErr)
		return
	}

	cards, appErr := h.paymentCardService.GetCardsByUserID(userID)
	if appErr != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		http2.AbortWithError(c, appErr)
		return
	}

	http2.RespondWithJSON(c, http.StatusOK,
		fromCardsToResponse(cards))
}

func (h *Handler) AddUserCards(c *gin.Context) {
	var req []CreateRequest
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		http2.AbortWithError(c, appErr)
		return
	}

	if err := c.BindJSON(&req); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		http2.AbortWithError(c, appErr)
		return
	}

	cards, appErr := h.paymentCardService.AddUserPaymentCard(userID,
		fromCreateRequestToCards(req))
	if appErr != nil {
		http2.AbortWithError(c, appErr)
		return
	}

	http2.RespondWithJSON(c, http.StatusCreated,
		fromCardsToResponse(cards))
}

func NewPaymentCardHandler(paymentCardService cardmanager.Manager) *Handler {
	return &Handler{paymentCardService}
}
