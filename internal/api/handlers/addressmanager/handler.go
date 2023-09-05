package addressmanager

import (
	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/usecases/addressmanager"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"
	"strconv"
)

func findAddressByUserId(m addressmanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		addresses, appErr := m.GetAddressesByUserId(userId)
		if appErr != nil {
			return appErr
		}

		response := fromAddressesToResponse(addresses)
		return httpUtil.RespondWithJSON(c, http.StatusOK, response)
	}
}

func registerAddressByUserId(m addressmanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request AddressCreate
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		if err = c.Bind(&request); err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		if err = request.Validate(); err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		address, appErr := m.AddUserAddress(userId, fromRequestToAddress(request))
		if appErr != nil {
			return appErr
		}

		response := fromAddressToResponse(address)
		return httpUtil.RespondWithJSON(c, http.StatusCreated, response)
	}
}
