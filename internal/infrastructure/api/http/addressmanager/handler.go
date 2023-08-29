package addressmanager

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/addressmanager"
)

type Handler struct {
	addressManager addressmanager.Manager
}

func Router(group *echo.Group) func(addressManagerHandler *Handler) {
	return func(handler *Handler) {
		group.GET("/address/find/:userId", handler.GetAddressesByUserId)
		group.POST("/address/register/:userId", handler.AddAddress)
	}
}

func (handler *Handler) GetAddressesByUserId(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	addresses, appErr := handler.addressManager.GetAddressesByUserId(userId)
	if appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusOK, fromAddressesToResponse(addresses))
}

func (handler *Handler) AddAddress(c echo.Context) error {
	var req RegisterAddressRequest
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	if err = c.Bind(&req); err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	address, appErr := handler.addressManager.AddUserAddress(userId,
		fromRequestToAddress(req))
	if appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusCreated, fromAddressToResponse(address))
}

func New(addressService addressmanager.Manager) *Handler {
	return &Handler{addressService}
}
