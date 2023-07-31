package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/constants"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/infra/api/dto"
	"github.com/rierarizzo/cafelatte/internal/infra/api/mappers"
	entities2 "github.com/rierarizzo/cafelatte/internal/infra/security/entities"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"net/http"
)

type AddressHandler struct {
	addressService ports.IAddressService
}

func (h *AddressHandler) AddUserAddresses(c *gin.Context) {
	var addressesRequest []dto.AddressRequest
	err := c.BindJSON(&addressesRequest)
	if err != nil {
		utils.AbortWithError(c, err)
		return
	}

	addresses := make([]entities.Address, 0)
	for _, v := range addressesRequest {
		addresses = append(addresses, *mappers.FromAddressReqToAddress(v))
	}

	userClaims := c.MustGet(constants.UserClaimsKey).(*entities2.UserClaims)

	addresses, err = h.addressService.AddUserAddresses(userClaims.ID, addresses)
	if err != nil {
		utils.AbortWithError(c, err)
		return
	}

	addressesRes := make([]dto.AddressResponse, 0)
	for _, v := range addresses {
		addressesRes = append(addressesRes, *mappers.FromAddressToAddressRes(v))
	}

	c.JSON(http.StatusCreated, addressesRes)
}

func NewAddressHandler(addressService ports.IAddressService) *AddressHandler {
	return &AddressHandler{addressService}
}
