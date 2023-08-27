package productpurchaser

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/productpurchaser"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/authenticator"
	"net/http"
)

type Handler struct {
	purchaser productpurchaser.Purchaser
}

func Router(group *echo.Group) func(purchaserHandler *Handler) {
	return func(handler *Handler) {
		group.Use(authenticator.Middleware)

		group.POST("/", handler.Purchase)
	}
}

func (handler *Handler) Purchase(c echo.Context) error {
	var req CreateOrderRequest
	if err := c.Bind(&req); err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	orderId, appErr := handler.purchaser.Purchase(fromRequestToOrder(req))
	if appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusCreated, fmt.Sprintf("Order with id %v created", orderId))
}

func New(purchaseUsecase productpurchaser.Purchaser) *Handler {
	return &Handler{purchaseUsecase}
}
