package cardmanager

import (
	"github.com/rierarizzo/cafelatte/internal/usecases/cardmanager"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func getCardsByUserId(m cardmanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		cards, appErr := m.GetCardsByUserId(userId)
		if appErr != nil {
			return appErr
		}

		response := fromCardsToResponse(cards)
		return httpUtil.RespondWithJSON(c, http.StatusOK, response)
	}
}

func addNewCard(m cardmanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		var request CardCreate
		if err = c.Bind(&request); err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		if err = request.Validate(); err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		card, appErr := m.AddUserCard(userId, fromRequestToCard(request))
		if appErr != nil {
			return appErr
		}

		response := fromCardToResponse(card)
		return httpUtil.RespondWithJSON(c, http.StatusCreated, response)
	}
}
