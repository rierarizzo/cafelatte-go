package handlers

import (
	"github.com/gin-gonic/gin"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/mappers"
	"github.com/rierarizzo/cafelatte/pkg/utils"
	"net/http"
	"strconv"
)

type PaymentCardHandler struct {
	paymentCardService ports.IPaymentCardService
}

func (h *PaymentCardHandler) GetCardsByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		utils.AbortWithError(c, appErr)
		return
	}

	cards, appErr := h.paymentCardService.GetCardsByUserID(userID)
	if appErr != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		utils.AbortWithError(c, appErr)
		return
	}

	utils.RespondWithJSON(c, http.StatusOK,
		mappers.FromCardSliceToCardResSlice(cards))
}

func (h *PaymentCardHandler) AddUserCards(c *gin.Context) {
	var req []dto.PaymentCardRequest
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		utils.AbortWithError(c, appErr)
		return
	}

	if err := c.BindJSON(&req); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		utils.AbortWithError(c, appErr)
		return
	}

	cards, appErr := h.paymentCardService.AddUserPaymentCard(userID,
		mappers.FromCardReqSliceToCardSlice(req))
	if appErr != nil {
		utils.AbortWithError(c, appErr)
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated,
		mappers.FromCardSliceToCardResSlice(cards))
}

func NewPaymentCardHandler(paymentCardService ports.IPaymentCardService) *PaymentCardHandler {
	return &PaymentCardHandler{paymentCardService}
}
