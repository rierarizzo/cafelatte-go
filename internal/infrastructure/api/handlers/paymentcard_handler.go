package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/constants"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/mappers"
	entities2 "github.com/rierarizzo/cafelatte/internal/infrastructure/security/claims"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"net/http"
)

type PaymentCardHandler struct {
	paymentCardService ports.IPaymentCardService
}

func (h *PaymentCardHandler) AddUserPaymentCards(c *gin.Context) {
	var cardsRequest []dto.PaymentCardRequest
	err := c.BindJSON(&cardsRequest)
	if err != nil {
		utils.AbortWithError(c, domain.NewAppError(err, domain.BadRequestError))
		return
	}

	cards := make([]entities.PaymentCard, 0)
	for _, v := range cardsRequest {
		cards = append(cards, mappers.FromPaymentCardReqToPaymentCard(v))
	}

	userClaims := c.MustGet(constants.UserClaimsKey).(*entities2.UserClaims)

	cards, appErr := h.paymentCardService.AddUserPaymentCard(userClaims.ID,
		cards)
	if appErr != nil {
		utils.AbortWithError(c, appErr)
		return
	}

	cardsRes := make([]dto.PaymentCardResponse, 0)
	for _, v := range cards {
		cardsRes = append(cardsRes, mappers.FromPaymentCardToPaymentCardRes(v))
	}

	c.JSON(http.StatusCreated, cardsRes)
}

func NewPaymentCardHandler(paymentCardService ports.IPaymentCardService) *PaymentCardHandler {
	return &PaymentCardHandler{paymentCardService}
}
