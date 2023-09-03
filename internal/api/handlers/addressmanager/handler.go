package addressmanager

import (
	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/usecases/addressmanager"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"
	"strconv"
)

func ConfigureRouting(g *echo.Group) func(m addressmanager.Manager) {
	addressGroup := g.Group("/address")

	return func(m addressmanager.Manager) {
		addressGroup.GET("/find/:userId", findAddressByUserId(m))
		addressGroup.POST("/register/:userId", registerAddressByUserId(m))
	}
}

func findAddressByUserId(m addressmanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
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
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		if err = c.Bind(&request); err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		if err = request.Validate(); err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		address, appErr := m.AddUserAddress(
			userId, fromRequestToAddress(request),
		)
		if appErr != nil {
			return appErr
		}

		response := fromAddressToResponse(address)
		return httpUtil.RespondWithJSON(c, http.StatusCreated, response)
	}
}
