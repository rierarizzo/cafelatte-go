package purchase

import (
	"github.com/gin-gonic/gin"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/purchase"
	http2 "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"
)

type Handler struct {
	purchaseUsecase purchase.IPurchaseUsecase
}

func (h *Handler) Purchase(c *gin.Context) {
	var req OrderRequest
	if err := c.BindJSON(&req); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		http2.AbortWithError(c, appErr)
		return
	}

	order := RequestToOrder(req)

	orderID, appErr := h.purchaseUsecase.PurchaseOrder(order)
	if appErr != nil {
		http2.AbortWithError(c, appErr)
		return
	}

	http2.RespondWithJSON(c, http.StatusCreated, gin.H{"orderID": orderID})
}

func NewPurchaseHandler(purchaseUsecase purchase.IPurchaseUsecase) *Handler {
	return &Handler{purchaseUsecase}
}
