package cardmanager

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/cardmanager"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"
	"strconv"
)

type Handler struct {
	paymentCardService cardmanager.Manager
}

func (h *Handler) GetCardsByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		httpUtil.AbortWithError(c, domain.NewAppError(err, domain.BadRequestError))
		return
	}

	cards, appErr := h.paymentCardService.GetCardsByUserID(userID)
	if appErr != nil {
		httpUtil.AbortWithError(c, appErr)
		return
	}

	httpUtil.RespondWithJSON(c, http.StatusOK, fromCardsToResponse(cards))
}

func (h *Handler) AddUserCards(c *gin.Context) {
	var req []RegisterCardRequest
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		httpUtil.AbortWithError(c, domain.NewAppError(err, domain.BadRequestError))
		return
	}

	if err := c.BindJSON(&req); err != nil {
		httpUtil.AbortWithError(c, domain.NewAppError(err, domain.BadRequestError))
		return
	}

	cards, appErr := h.paymentCardService.AddUserPaymentCard(userID,
		fromRequestToCards(req))
	if appErr != nil {
		httpUtil.AbortWithError(c, appErr)
		return
	}

	httpUtil.RespondWithJSON(c, http.StatusCreated, fromCardsToResponse(cards))
}

func NewPaymentCardHandler(paymentCardService cardmanager.Manager) *Handler {
	return &Handler{paymentCardService}
}
