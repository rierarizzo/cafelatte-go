package handlers

import (
	"github.com/gin-gonic/gin"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/mappers"
	"github.com/rierarizzo/cafelatte/pkg/utils"
	"net/http"
)

type PurchaseHandler struct {
	purchaseUsecase ports.IPurchaseUsecase
}

func (h *PurchaseHandler) Purchase(c *gin.Context) {
	var req dto.PurchaseOrderRequest
	if err := c.BindJSON(&req); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		utils.AbortWithError(c, appErr)
		return
	}

	order := mappers.RequestToOrder(req)

	orderID, appErr := h.purchaseUsecase.PurchaseOrder(order)
	if appErr != nil {
		utils.AbortWithError(c, appErr)
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated, gin.H{"orderID": orderID})
}

func NewPurchaseHandler(purchaseUsecase ports.IPurchaseUsecase) *PurchaseHandler {
	return &PurchaseHandler{purchaseUsecase}
}
