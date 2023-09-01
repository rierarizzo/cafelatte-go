package productpurchaser

import (
	"fmt"
	"github.com/rierarizzo/cafelatte/internal/usecases/productpurchaser"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
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

		response := fmt.Sprintf("Order with id %v created", orderId)
		return httpUtil.RespondWithJSON(c, http.StatusCreated, response)
	}
}
