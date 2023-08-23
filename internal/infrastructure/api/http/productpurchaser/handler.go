package productpurchaser

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/productpurchaser"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"
)

type Handler struct {
	purchaser productpurchaser.Purchaser
}

func (h *Handler) Purchase(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.BindJSON(&req); err != nil {
		httpUtil.AbortWithError(c, domain.NewAppError(err, domain.BadRequestError))
		return
	}

	orderID, appErr := h.purchaser.Purchase(fromRequestToOrder(req))
	if appErr != nil {
		httpUtil.AbortWithError(c, appErr)
		return
	}

	httpUtil.RespondWithJSON(c, http.StatusCreated, gin.H{"orderID": orderID})
}

func NewPurchaseHandler(purchaseUsecase productpurchaser.Purchaser) *Handler {
	return &Handler{purchaseUsecase}
}
