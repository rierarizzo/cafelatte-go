package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/constants"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/infra/api/dto"
	"github.com/rierarizzo/cafelatte/internal/infra/api/mappers"
	entities2 "github.com/rierarizzo/cafelatte/internal/infra/security/entities"
	"log/slog"
	"net/http"
)

type PaymentCardHandler struct {
	paymentCardService ports.IPaymentCardService
}

func (h *PaymentCardHandler) AddUserPaymentCards(c *gin.Context) {
	var cardsRequest []dto.PaymentCardRequest
	err := c.BindJSON(&cardsRequest)
	if err != nil {
		slog.Error(c.Error(err).Error())
		c.Abort()
		return
	}

	cards := make([]entities.PaymentCard, 0)
	for _, v := range cardsRequest {
		cards = append(cards, *mappers.FromPaymentCardReqToPaymentCard(v))
	}

	userClaims := c.MustGet(constants.UserClaimsKey).(*entities2.UserClaims)

	cards, err = h.paymentCardService.AddUserPaymentCard(userClaims.ID, cards)
	if err != nil {
		slog.Error(c.Error(err).Error())
		c.Abort()
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
