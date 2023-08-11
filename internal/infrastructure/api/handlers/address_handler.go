package handlers

import (
	"github.com/gin-gonic/gin"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/mappers"
	"github.com/rierarizzo/cafelatte/pkg/utils"
	"net/http"
	"strconv"
)

type AddressHandler struct {
	addressService ports.IAddressService
}

func (h *AddressHandler) GetAddressesByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		utils.AbortWithError(c, appErr)
		return
	}

	addresses, appErr := h.addressService.GetAddressesByUserID(userID)
	if appErr != nil {
		utils.AbortWithError(c, appErr)
		return
	}

	utils.RespondWithJSON(c, http.StatusOK,
		mappers.FromAddressSliceToAddressResSlice(addresses))
}

func (h *AddressHandler) AddUserAddresses(c *gin.Context) {
	var req []dto.AddressRequest
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		utils.AbortWithError(c, appErr)
		return
	}

	if err := c.BindJSON(&req); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		utils.AbortWithError(c, appErr)
		return
	}

	addresses, appErr := h.addressService.AddUserAddresses(userID,
		mappers.FromAddressReqSliceToAddressSlice(req))
	if appErr != nil {
		utils.AbortWithError(c, appErr)
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated,
		mappers.FromAddressSliceToAddressResSlice(addresses))
}

func NewAddressHandler(addressService ports.IAddressService) *AddressHandler {
	return &AddressHandler{addressService}
}
