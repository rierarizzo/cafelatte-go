package productpurchaser

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/productpurchaser"
)

type Handler struct {
	purchaser productpurchaser.Purchaser
}

func Router(group *echo.Group) func(purchaserHandler *Handler) {
	return func(h *Handler) {
		group.POST("/", h.Purchase)
	}
}

func (h *Handler) Purchase(c echo.Context) error {
	var req OrderCreate
	if err := c.Bind(&req); err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	orderId, appErr := h.purchaser.Purchase(fromRequestToOrder(req))
	if appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusCreated,
		fmt.Sprintf("Order with id %v created", orderId))
}

func New(productPurchaser productpurchaser.Purchaser) *Handler {
	return &Handler{productPurchaser}
}
