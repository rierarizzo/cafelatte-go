package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/core/constants"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
	errorHandler "github.com/rierarizzo/cafelatte/internal/infrastructure/api/error"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/mappers"
	"net/http"
)

type PaymentCardHandler struct {
	paymentCardService ports.IPaymentCardService
}

func (h *PaymentCardHandler) AddUserPaymentCards(c *gin.Context) {
	var cardsRequest []dto.PaymentCardRequest
	err := c.BindJSON(&cardsRequest)
	if errorHandler.Error(c, err) {
		return
	}

	cards := make([]entities.PaymentCard, 0)
	for _, v := range cardsRequest {
		cards = append(cards, *mappers.FromPaymentCardReqToPaymentCard(v))
	}

	userClaims := c.MustGet(constants.UserClaimsKey).(*entities.UserClaims)

	cards, err = h.paymentCardService.AddUserPaymentCard(userClaims.ID, cards)
	if errorHandler.Error(c, err) {
		return
	}

	cardsRes := make([]dto.PaymentCardResponse, 0)
	for _, v := range cards {
		cardsRes = append(cardsRes, *mappers.FromPaymentCardToPaymentCardRes(v))
	}

	c.JSON(http.StatusCreated, cardsRes)
}

func NewPaymentCardHandler(paymentCardService ports.IPaymentCardService) *PaymentCardHandler {
	return &PaymentCardHandler{paymentCardService}
}
