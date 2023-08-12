package address

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/address"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"
	"strconv"
)

type Handler struct {
	addressService address.IAddressService
}

func (h *Handler) GetAddressesByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		httpUtil.AbortWithError(c, appErr)
		return
	}

	addresses, appErr := h.addressService.GetAddressesByUserID(userID)
	if appErr != nil {
		httpUtil.AbortWithError(c, appErr)
		return
	}

	httpUtil.RespondWithJSON(c, http.StatusOK,
		fromAddressesToResponse(addresses))
}

func (h *Handler) AddUserAddresses(c *gin.Context) {
	var req []CreateRequest
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		httpUtil.AbortWithError(c, appErr)
		return
	}

	if err := c.BindJSON(&req); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		httpUtil.AbortWithError(c, appErr)
		return
	}

	addresses, appErr := h.addressService.AddUserAddresses(userID,
		fromCreateRequestToAddresses(req))
	if appErr != nil {
		httpUtil.AbortWithError(c, appErr)
		return
	}

	httpUtil.RespondWithJSON(c, http.StatusCreated,
		fromAddressesToResponse(addresses))
}

func NewAddressHandler(addressService address.IAddressService) *Handler {
	return &Handler{addressService}
}
