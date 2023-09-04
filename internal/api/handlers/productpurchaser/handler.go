package productpurchaser

import (
	"fmt"
	"github.com/rierarizzo/cafelatte/internal/usecases/productpurchaser"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func purchase(p productpurchaser.Purchaser) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request OrderCreate
		if err := c.Bind(&request); err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		orderId, appErr := p.Purchase(fromRequestToOrder(request))
		if appErr != nil {
			return appErr
		}

		response := fmt.Sprintf("Order with id %v created", orderId)
		return httpUtil.RespondWithJSON(c, http.StatusCreated, response)
	}
}
