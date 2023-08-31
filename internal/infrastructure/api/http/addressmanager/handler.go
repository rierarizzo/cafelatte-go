package addressmanager

import (
	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/addressmanager"
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
		return c.JSON(http.StatusOK, response)
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

		err = c.Bind(&request)
		if err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		address, appErr := m.AddUserAddress(userId,
			fromRequestToAddress(request))
		if appErr != nil {
			return appErr
		}

		response := fromAddressToResponse(address)
		return c.JSON(http.StatusCreated, response)
	}
}
