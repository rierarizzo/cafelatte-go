package productpurchaser

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/productpurchaser"
	http2 "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"
)

type Handler struct {
	purchaser productpurchaser.Purchaser
}

func (h *Handler) Purchase(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.BindJSON(&req); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		http2.AbortWithError(c, appErr)
		return
	}

	order := fromCreateOrderRequestToOrder(req)

	orderID, appErr := h.purchaser.Purchase(order)
	if appErr != nil {
		http2.AbortWithError(c, appErr)
		return
	}

	http2.RespondWithJSON(c, http.StatusCreated, gin.H{"orderID": orderID})
}

func NewPurchaseHandler(purchaseUsecase productpurchaser.Purchaser) *Handler {
	return &Handler{purchaseUsecase}
}
