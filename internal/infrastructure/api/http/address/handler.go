package address

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/address"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/user"
	http2 "github.com/rierarizzo/cafelatte/pkg/utils/http"
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
		http2.AbortWithError(c, appErr)
		return
	}

	addresses, appErr := h.addressService.GetAddressesByUserID(userID)
	if appErr != nil {
		http2.AbortWithError(c, appErr)
		return
	}

	http2.RespondWithJSON(c, http.StatusOK,
		FromAddressSliceToAddressResSlice(addresses))
}

func (h *Handler) AddUserAddresses(c *gin.Context) {
	var req []user.AddressRequest
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		http2.AbortWithError(c, appErr)
		return
	}

	if err := c.BindJSON(&req); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		http2.AbortWithError(c, appErr)
		return
	}

	addresses, appErr := h.addressService.AddUserAddresses(userID,
		FromAddressReqSliceToAddressSlice(req))
	if appErr != nil {
		http2.AbortWithError(c, appErr)
		return
	}

	http2.RespondWithJSON(c, http.StatusCreated,
		FromAddressSliceToAddressResSlice(addresses))
}

func NewAddressHandler(addressService address.IAddressService) *Handler {
	return &Handler{addressService}
}
