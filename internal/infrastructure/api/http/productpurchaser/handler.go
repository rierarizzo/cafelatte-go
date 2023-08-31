package productpurchaser

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/productpurchaser"
)

func ConfigureRouting(g *echo.Group) func(p productpurchaser.Purchaser) {
	return func(p productpurchaser.Purchaser) {
		g.POST("/", purchase(p))
	}
}

func purchase(p productpurchaser.Purchaser) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request OrderCreate
		err := c.Bind(&request)
		if err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		orderId, appErr := p.Purchase(fromRequestToOrder(request))
		if appErr != nil {
			return appErr
		}

		return c.JSON(http.StatusCreated,
			fmt.Sprintf("Order with id %v created", orderId))
	}
}
